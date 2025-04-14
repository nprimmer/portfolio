package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	PageSize     = 1024
	SearchString = "gc24{REDACTED}"
	MinPages     = 1
	MaxPages     = 15
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number of pages between MinPages and MaxPages
	totalPages := rand.Intn(MaxPages-MinPages+1) + MinPages

	file, err := os.Create("example.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < totalPages; i++ {
		page := make([]byte, PageSize)
		if i == 7 {
			// Insert the search string into page 500
			copy(page, []byte(SearchString))
		} else {
			// Fill the page with random characters
			for j := range page {
				page[j] = byte(rand.Intn(256))
			}
		}
		file.Write(page)
	}

	fmt.Println("Large file created successfully with", totalPages, "pages.")
}
