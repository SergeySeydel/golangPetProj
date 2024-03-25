package main

import (
	"fmt"
	"strings"
)

func main(){

	digits := "0123456789"
	lowerAlph := "abcdefghijklmnopqrstuvwxyz"
	upperAlph := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special := "_!@#$%^&"
	maxAttempts := 5

	for attempt := 1; attempt <= maxAttempts; attempt++{
		var password string
		fmt.Printf("Try number %d. Please, enter password: ", attempt)
		fmt.Scanln(&password)

		if len(password) < 8 || len(password) > 15{
			fmt.Printf("Password must contain from 8 to 15 characters")
			continue
		}
		
		var (
			hasDigit bool
			hasLower bool
			hasUpper bool
			hasSpecial bool
		)

		for _, char := range password{
			switch{
			case strings.ContainsRune(digits, char):
				hasDigit = true
			case strings.ContainsRune(lowerAlph, char):
				hasLower = true
			case strings.ContainsRune(upperAlph, char):
				hasUpper = true
			case strings.ContainsRune(special, char):
				hasSpecial = true
			}
		}
		if !hasDigit || !hasLower || !hasUpper || !hasSpecial{
			fmt.Println("Password must contain at least one symbol from this set: digit, lower register character, upper register character, special character")
			continue
		}
		fmt.Println("You have entered right password")
		return
	}

	fmt.Println("Maximum number of attempts to enter valid password has been exceeded")
	return
}