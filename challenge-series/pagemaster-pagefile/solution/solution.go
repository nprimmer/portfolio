package main

import (
	"errors"
)

// This file contains an example of the solution to this challenge.
// In this implementation, the function pages through the pagefile,
// until it encounters the string "gc24". Upon locating the substring,
// The application reads to the ending brace.
func SearchString(mem *Memory) (string, error) {
	searchPrefix := "gc24"
	buffer := make([]byte, len(searchPrefix))
	foundAt := -1
	totalAddresses := TotalPages * PageSize

	// Search for the prefix "gc24"
	for address := 0; address < totalAddresses; address++ {
		byteVal, err := mem.ReadAddress(address)
		if err != nil {
			return "", err
		}

		buffer = append(buffer[1:], byteVal)
		if string(buffer) == searchPrefix {
			foundAt = address - len(searchPrefix) + 1
			break
		}
	}

	if foundAt == -1 {
		return "", errors.New("string not found")
	}

	// Read the full string wrapped in "gc24{}"
	fullString := searchPrefix
	for i := foundAt + len(searchPrefix); ; i++ {
		byteVal, err := mem.ReadAddress(i)
		if err != nil {
			return "", err
		}
		fullString += string(byteVal)
		if byteVal == '}' {
			break
		}
	}

	return fullString, nil
}
