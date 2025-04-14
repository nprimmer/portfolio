package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go/ast"
	"go/parser"
	"go/token"

	"github.com/gin-gonic/gin"
)

var (
	localFilePath = getEnv("LOCAL_FILE_PATH", "/app/memory.go")
	pageFilePath  = getEnv("PAGE_FILE_PATH", "/app/pagefile.dat")
)

func main() {
	log.Println("Starting server on :8080")
	r := gin.Default()
	r.POST("/validate", validateURL)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func validateURL(c *gin.Context) {
	var json struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.BindJSON(&json); err != nil {
		log.Printf("Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	url := json.URL
	log.Printf("Received URL: %s", url)
	if strings.HasSuffix(url, ".go") {
		handleRawGoFile(url, c)
	} else if strings.Contains(url, "github.com") {
		handleGitHubRepo(url, c)
	} else {
		log.Printf("Provided URL does not resolve: %s", url)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Provided URL does not resolve"})
	}
}

func handleGitHubRepo(url string, c *gin.Context) {
	dir, err := os.MkdirTemp("", "repo")
	if err != nil {
		log.Printf("Failed to create temp directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temp directory"})
		return
	}

	log.Printf("Cloning repository: %s", url)
	cloneCmd := exec.Command("git", "clone", url, dir)
	if err := cloneCmd.Run(); err != nil {
		log.Printf("Failed to clone repository: %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Provided URL does not resolve"})
		return
	}

	mainFilePath := filepath.Join(dir, "main.go")
	handleGoFile(mainFilePath, dir, c)

	// Schedule directory deletion after 30 seconds
	go scheduleTempDirDeletion(dir)
}

func handleRawGoFile(url string, c *gin.Context) {
	log.Printf("Fetching raw Go file: %s", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Failed to fetch the file: %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Provided URL does not resolve"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read the file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read the file"})
		return
	}

	dir, err := os.MkdirTemp("", "go-file")
	if err != nil {
		log.Printf("Failed to create temp directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temp directory"})
		return
	}

	filePath := filepath.Join(dir, "main.go")
	if err := ioutil.WriteFile(filePath, body, 0644); err != nil {
		log.Printf("Failed to write the file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write the file"})
		return
	}

	handleGoFile(filePath, dir, c)

	// Schedule directory deletion after 30 seconds
	go scheduleTempDirDeletion(dir)
}

func handleGoFile(filePath string, dir string, c *gin.Context) {
	log.Printf("Validating Go file: %s", filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read the file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read the file"})
		return
	}

	if err := validateGoFile(content, c); err != nil {
		log.Printf("Validation failed: %v", err)
		return
	}

	// Copy local file to temp directory
	log.Printf("Copying local file: %s", localFilePath)
	localContent, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		log.Printf("Failed to read local file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read local file"})
		return
	}
	localFileCopyPath := filepath.Join(dir, filepath.Base(localFilePath))
	if err := ioutil.WriteFile(localFileCopyPath, localContent, 0644); err != nil {
		log.Printf("Failed to write local file to temp directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write local file to temp directory"})
		return
	}

	// Copy page file to temp directory
	log.Printf("Copying page file: %s", pageFilePath)
	pageFileContent, err := ioutil.ReadFile(pageFilePath)
	if err != nil {
		log.Printf("Failed to read page file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read page file"})
		return
	}
	pageFileCopyPath := filepath.Join(dir, filepath.Base(pageFilePath))
	if err := ioutil.WriteFile(pageFileCopyPath, pageFileContent, 0644); err != nil {
		log.Printf("Failed to write page file to temp directory: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write page file to temp directory"})
		return
	}

	// Run `go run main.go memory.go` in the temp directory
	log.Printf("Running `go run main.go memory.go` in directory: %s", dir)
	cmd := exec.Command("go", "run", "main.go", "memory.go")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Compilation failed: %v\nOutput: %s", err, string(output))
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Provided file failed to compile", "details": string(output)})
		return
	}

	log.Printf("Compilation succeeded with output: %s", string(output))
	c.JSON(http.StatusOK, gin.H{"output": string(output)})
}

func validateGoFile(content []byte, c *gin.Context) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", bytes.NewReader(content), parser.AllErrors)
	if err != nil {
		log.Printf("Not a valid Go file: %v", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Not a valid Go file"})
		return err
	}

	if node.Name.Name != "main" {
		log.Printf("File is not package main")
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "File must be package main"})
		return err
	}

	found := false
	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if funcDecl.Name.Name == "SearchString" {
				if len(funcDecl.Type.Params.List) == 1 && len(funcDecl.Type.Results.List) == 2 {
					if funcDecl.Type.Params.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name == "Memory" &&
						funcDecl.Type.Results.List[0].Type.(*ast.Ident).Name == "string" &&
						funcDecl.Type.Results.List[1].Type.(*ast.Ident).Name == "error" {
						found = true
						break
					}
				}
			}
		}
	}

	if !found {
		log.Printf("Provided file does not match the expected signature")
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Provided file does not match the expected signature"})
		return err
	}

	return nil
}

func scheduleTempDirDeletion(dir string) {
	log.Printf("Scheduled deletion of temp directory: %s", dir)
	time.Sleep(30 * time.Second)
	if err := os.RemoveAll(dir); err != nil {
		log.Printf("Failed to delete temp directory: %v", err)
	} else {
		log.Printf("Temp directory deleted: %s", dir)
	}
}
