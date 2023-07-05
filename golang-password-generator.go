//Main entry to the program
package main

import (
	//Import strong random number generation tool
	"crypto/rand"
	//Import format functions for formatted output 
	"fmt"
)

//Constant values
const (
	//Password character sets to use. Identifier = value. 
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+-=[]{}|;:,.<>?/~"
)
