package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/pbkdf2"
)

type Flag struct {
	Flag string `json:"flag"`
}

type EncryptedFlag struct {
	FlagEnc string `json:"flag"`
}

var db *sql.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var err error
	db, err = sql.Open("sqlite3", "./id_disc.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable()

	r := gin.Default()
	r.Use(authMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "200 OK")
	})

	r.POST("/MakeFlag", makeFlagHandler)
	r.GET("/GetFlags", getFlagsHandler)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS flags (
		flag_checksum TEXT PRIMARY KEY,
		key_checksum TEXT,
		flag_enc TEXT
	);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth {
			c.Header("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authStr := fmt.Sprintf("%s:%s", username, password)
		key := pbkdf2.Key([]byte(authStr), []byte("some_salt"), 4096, 32, sha256.New)
		encKey := base64.StdEncoding.EncodeToString(key)

		c.Set("encKey", encKey)
		c.Next()
	}
}

func makeFlagHandler(c *gin.Context) {
	var flag Flag
	if err := c.ShouldBindJSON(&flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encKey, _ := c.Get("encKey")
	encKeyStr := encKey.(string)

	flagEnc := encrypt(encKeyStr, flag.Flag)
	flagChecksum := fmt.Sprintf("%x", sha1.Sum([]byte(flag.Flag)))
	keyChecksum := fmt.Sprintf("%x", sha1.Sum([]byte(encKeyStr)))

	query := `INSERT INTO flags (flag_checksum, key_checksum, flag_enc) VALUES (?, ?, ?)`
	if _, err := db.Exec(query, flagChecksum, keyChecksum, flagEnc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, EncryptedFlag{FlagEnc: flagEnc})
}

func getFlagsHandler(c *gin.Context) {
	encKey, _ := c.Get("encKey")
	encKeyStr := encKey.(string)

	keyChecksum := fmt.Sprintf("%x", sha1.Sum([]byte(encKeyStr)))

	query := `SELECT flag_enc FROM flags WHERE key_checksum = ?`
	rows, err := db.Query(query, keyChecksum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	flags := make(map[string]string)
	i := 0
	for rows.Next() {
		var flagEnc string
		if err := rows.Scan(&flagEnc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		decryptedFlag := decrypt(encKeyStr, flagEnc)
		flags[fmt.Sprintf("flag_%d", i)] = decryptedFlag
		i++
	}

	c.JSON(http.StatusOK, flags)
}

func encrypt(key string, text string) string {
	// Use a simple XOR encryption for demonstration purposes
	encrypted := make([]byte, len(text))
	for i := range text {
		encrypted[i] = text[i] ^ key[i%len(key)]
	}
	return base64.StdEncoding.EncodeToString(encrypted)
}

func decrypt(key string, cipherText string) string {
	data, _ := base64.StdEncoding.DecodeString(cipherText)
	decrypted := make([]byte, len(data))
	for i := range data {
		decrypted[i] = data[i] ^ key[i%len(key)]
	}
	return string(decrypted)
}
