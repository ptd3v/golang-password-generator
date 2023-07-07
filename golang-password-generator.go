// Main entry to the program
package main

import (
	//A random number generator.
	"math/rand"
	//An alterative for a more secure random number generator would be "crypto/rand".
	//crypto/rand
	//Provides format functions for formatted output e.g fmt.Println
	"fmt"
)

// Constant values
const (
	//Password character sets to use. Identifier = value.
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+-=[]{}|;:,.<>?/~"
)

// Create a blank password string
func generatePassword() string {
	//Concatenate all characters into one unified charset
	charset := uppercaseLetters + lowercaseLetters + digits + specialChars
	//Initialize an empty string to store the generated password
	password := ""

	//Generate a password of length 20
	for i := 0; i < 20; i++ {
		//Randomly select a character from the charset and append it to the password string
		password += string(charset[rand.Intn(len(charset))])
	}

	//Return the generated password to the string
	return password
}

func main() {
	//Fetch the newly created password string
	password := generatePassword()
	//Print the newly generated password
	fmt.Println("Generated Password:", password)
}
