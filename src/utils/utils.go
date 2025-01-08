package utils

import (
	"fmt"
)

// PrintMessage prints a message to the console
func PrintMessage(message string) {
	fmt.Println(message)
}

// LogError logs an error message
func LogError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
