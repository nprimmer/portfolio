package main

import (
	"fmt"
	"os"
)

func main() {
	// Example usage
	file, err := os.Open("../utils/pagefile.dat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mem := NewMemory(file)

	foundString, err := SearchString(mem)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Print(foundString)
	}
}
