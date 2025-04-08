package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var mostCommonPasswords = []string{
	"123456",
	"123456789",
	"12345",
	"qwerty",
	"password",
	"12345678",
	"111111",
	"123123",
	"qwerty123",
	"000000",
	"1q2w3e",
	"aa12345678",
	"abc123",
	"password1",
	"1234",
	"qwertyuiop",
	"123321",
	"password123",
}

type User struct {
	gorm.Model
	Username string `json:"user" binding:"required" gorm:"uniqueIndex"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Admin    bool   `json:"admin"`
	Title    string `json:"title"`
}

var db *gorm.DB
var err error

var users = []User{
	{
		Username: "a.johnson",
		Email:    "a.johnson@nova-messages.mentats.org",
		Password: "vK8!zU4#pQeL",
		Admin:    true,
		Title:    "Head of IT",
	},
	{
		Username: "e.carter",
		Email:    "e.carter@nova-messages.mentats.org",
		Password: "yT3@hM7!cUzR",
		Admin:    true,
		Title:    "Head of Research",
	},
	{
		Username: "g.anderson",
		Email:    "g.anderson@nova-messages.mentats.org",
		Password: mostCommonPasswords[0],
		Admin:    false,
		Title:    "Senior Engineer",
	},
	{
		Username: "c.thompson",
		Email:    "c.thompson@nova-messages.mentats.org",
		Password: "nJ6@bK1!rWxZ",
		Admin:    false,
		Title:    "Finance Manager",
	},
	{
		Username: "d.lee",
		Email:    "d.lee@nova-messages.mentats.org",
		Password: "pF2@dL8!mTqK",
		Admin:    false,
		Title:    "HR Director",
	},
	{
		Username: "e.davis",
		Email:    "e.davis@nova-messages.mentats.org",
		Password: "gV7@rN2!tWxM",
		Admin:    false,
		Title:    "Lead Developer",
	},
	{
		Username: "f.wilson",
		Email:    "f.wilson@nova-messages.mentats.org",
		Password: "sM5@vB3!pQtR",
		Admin:    false,
		Title:    "Marketing Manager",
	},
	{
		Username: "g.martinez",
		Email:    "g.martinez@nova-messages.mentats.org",
		Password: "kL4@cV6!zRyP",
		Admin:    false,
		Title:    "Administrative Assistant",
	},
	{
		Username: "h.brown",
		Email:    "h.brown@nova-messages.mentats.org",
		Password: "mT8@dK2!yNzQ",
		Admin:    false,
		Title:    "Operations Manager",
	},
	{
		Username: "i.white",
		Email:    "i.white@nova-messages.mentats.org",
		Password: "fP9@lM7!tQrX",
		Admin:    false,
		Title:    "Quality Assurance Lead",
	},
	{
		Username: "j.green",
		Email:    "j.green@nova-messages.mentats.org",
		Password: "rQ2@vL5!nMwT",
		Admin:    false,
		Title:    "Robotics Engineer",
	},
	{
		Username: "k.roberts",
		Email:    "k.roberts@nova-messages.mentats.org",
		Password: "dM3@pN8!zLtQ",
		Admin:    false,
		Title:    "Data Analyst",
	},
	{
		Username: "l.king",
		Email:    "l.king@nova-messages.mentats.org",
		Password: "tK6@xP4!qMzV",
		Admin:    false,
		Title:    "Network Administrator",
	},
	{
		Username: "m.taylor",
		Email:    "m.taylor@nova-messages.mentats.org",
		Password: "cQ8@rL1!vPwY",
		Admin:    false,
		Title:    "Product Manager",
	},
	{
		Username: "n.scott",
		Email:    "n.scott@nova-messages.mentats.org",
		Password: "xV7@tM2!kPrL",
		Admin:    false,
		Title:    "Security Specialist",
	},
}

func userExists(username string) bool {
	var count int64
	db.Model(&User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func validateCredentials(username string, password string) bool {
	var count int64
	db.Model(&User{}).Where("username = ? AND password = ?", username, password).Count(&count)
	return count > 0
}

func resetGregPassword() {
	rand.Seed(time.Now().UnixNano())

	newPassword := mostCommonPasswords[rand.Intn(len(mostCommonPasswords))]

	db.Model(&User{}).Where("username = ?", "g.anderson").Update("password", newPassword)
}

func isAdminUser(username string) bool {
	var count int64
	db.Model(&User{}).Where("username = ? AND admin = true", username).Count(&count)
	return count > 0
}

func deleteUser(username string) {
	db.Where("username = ?", username).Delete(&User{})
}

func getFlag(c *gin.Context) {
	user, password, _ := c.Request.BasicAuth()
	if user == "" && password == "" {
		c.String(http.StatusUnauthorized, "Failed to provide credentials")
		return
	}
	if !validateCredentials(user, password) {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	c.String(http.StatusOK, "gc24{41ece812-ecc2-46d4-969d-3008a3d77e70}\n")
	c.String(http.StatusOK, "Suspicious activity detected! Administrator has been notified!")
	if user == "g.anderson" {
		c.String(http.StatusOK, "Suspicious activity detected! Administrator has been notified!")
		resetGregPassword()
	}
}

func getAdminFlag(c *gin.Context) {
	user, password, _ := c.Request.BasicAuth()
	if user == "" && password == "" {
		c.String(http.StatusUnauthorized, "Failed to provide credentials")
		return
	}
	if !validateCredentials(user, password) {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	if !isAdminUser(user) {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	c.String(http.StatusOK, "gc24{0a45ace0-cd56-434e-9dec-7053c0482236}\n")
	c.String(http.StatusOK, "Suspicious activity detected! Administrator has been notified!")
	deleteUser(user)
	c.String(http.StatusOK, "Greg is forever, but %s is temporary", user)
}

func createUser(c *gin.Context) {
	user, password, _ := c.Request.BasicAuth()
	if !validateCredentials(user, password) {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request."})
		return
	}

	if userExists(newUser.Username) {
		c.String(http.StatusConflict, "Invalid username.")
		return
	}
	if !newUser.Admin {
		newUser.Admin = true
	}
	db.Create(&newUser)
	c.String(http.StatusOK, "User created!")
	if user == "g.anderson" {
		c.String(http.StatusOK, "Suspicious activity detected! Administrator has been notified!")
		resetGregPassword()
	}
}

func getUser(c *gin.Context) {
	username := c.Param("username")

	var u User
	if err := db.Where("username = ?", username).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		return
	}

	// Redact the password
	u.Password = "REDACTED"

	c.JSON(http.StatusOK, u)
}

func getMessage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format."})
		return
	}

	for _, msg := range messages {
		if int(msg.ID) == id {
			c.JSON(http.StatusOK, msg)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found."})
}

func main() {
	log.Info("Starting...")
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to initialize database")
	}
	log.Infof("Database initialized")
	db.AutoMigrate(&User{})
	for _, u := range users {
		log.Infof("Creating User: %s", u.Username)
		db.Create(&u)
		log.Infof("User created")
	}
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://nova-messages.mentats.org"}
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	r.GET("/get-flag", getFlag)
	r.GET("/get-admin-flag", getAdminFlag)
	r.GET("/user/:username", getUser)
	r.GET("/message/:id", getMessage)
	r.POST("/user", createUser)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, "Invalid resource reference")
	})
	r.Run()
}
