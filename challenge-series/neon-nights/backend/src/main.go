package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// IssuedDate structure
type IssuedDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

// ExpirationDate structure
type ExpirationDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day,omitempty"`
}

// Address structure
type Address struct {
	Number     int    `json:"number"`
	Street     string `json:"street"`
	StreetType string `json:"street_type"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zip        int    `json:"zip"`
}

// DebitCard structure
type DebitCard struct {
	Type           string         `json:"type"`
	Name           string         `json:"name"`
	Number         string         `json:"number"`
	IssuedDate     IssuedDate     `json:"issued_date"`
	ExpirationDate ExpirationDate `json:"expiration_date"`
}

// CreditCard structure
type CreditCard struct {
	Type           string         `json:"type"`
	CardType       string         `json:"card_type"`
	Name           string         `json:"name"`
	Number         string         `json:"number"`
	IssuedDate     IssuedDate     `json:"issued_date"`
	ExpirationDate ExpirationDate `json:"expiration_date"`
}

// DriverLicense structure
type DriverLicense struct {
	Type           string         `json:"type"`
	Name           string         `json:"name"`
	Number         string         `json:"number"`
	ExpirationDate ExpirationDate `json:"expiration_date"`
	Address        Address        `json:"address"`
	Birthdate      ExpirationDate `json:"birthdate"`
}

// Note structure
type Note struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ATMReceipt structure
type ATMReceipt struct {
	Type          string `json:"type"`
	TransactionID string `json:"transaction_id"`
}

// WalletItem structure
type WalletItem struct {
	Type    string          `json:"type"`
	Details json.RawMessage `json:"details"`
}

// Person structure
type Person struct {
	WalletContents []WalletItem `json:"wallet_contents"`
	Pin            string       `json:"pin"`
	NextKey        string       `json:"next_key"`
}

var people map[string]Person

func loadPeople(filePath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &people)
}

func getWalletContents(walletItems []WalletItem) ([]interface{}, error) {
	var contents []interface{}
	for _, item := range walletItems {
		var detail interface{}
		switch item.Type {
		case "debit_card":
			var debitCard DebitCard
			if err := json.Unmarshal(item.Details, &debitCard); err != nil {
				return nil, err
			}
			debitCard.Type = "debit_card"
			detail = debitCard
		case "credit_card":
			var creditCard CreditCard
			if err := json.Unmarshal(item.Details, &creditCard); err != nil {
				return nil, err
			}
			creditCard.Type = "credit_card"
			detail = creditCard
		case "driver_license":
			var driverLicense DriverLicense
			if err := json.Unmarshal(item.Details, &driverLicense); err != nil {
				return nil, err
			}
			driverLicense.Type = "driver_license"
			detail = driverLicense
		case "note":
			var note Note
			if err := json.Unmarshal(item.Details, &note); err != nil {
				return nil, err
			}
			note.Type = "note"
			detail = note
		case "atm_receipt":
			var atmReceipt ATMReceipt
			if err := json.Unmarshal(item.Details, &atmReceipt); err != nil {
				return nil, err
			}
			atmReceipt.Type = "atm_receipt"
			detail = atmReceipt
		default:
			continue
		}
		contents = append(contents, detail)
	}
	return contents, nil
}

func main() {
	err := loadPeople("people.json")
	if err != nil {
		panic("Failed to load people.json: " + err.Error())
	}

	r := gin.Default()

	// CORS configuration
	config := cors.Config{
		AllowOrigins:     []string{"https://miami-first-federal-of-miami.toolchest.app", "https://first-miami-backend.toolchest.app"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	// Dynamic endpoint creation
	for key, person := range people {
		person := person // capture range variable
		r.GET("/"+key, func(c *gin.Context) {
			contents, err := getWalletContents(person.WalletContents)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse wallet contents"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"wallet_contents": contents})
		})

		r.POST("/"+key, func(c *gin.Context) {
			var request struct {
				Pin string `json:"pin"`
			}
			if err := c.BindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
				return
			}

			if request.Pin == person.Pin {
				c.JSON(http.StatusOK, gin.H{"next_key": person.NextKey})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid PIN"})
			}
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run(":8080")
}
