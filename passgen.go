package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

func generatePassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[randomNum.Int64()]
	}
	return string(password), nil
}

func main() {
	var size int
	fmt.Println("Enter the password length (1 to 128): ")
	fmt.Scanf("%d", &size)
	if size < 1 || size > 128 {
		fmt.Printf("Invalid number %d %s\n", size, "program will be closed.")
		return
	}
	password, err := generatePassword(size)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println("Generated password:", password)
}
