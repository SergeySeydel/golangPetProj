package main

import "fmt"

/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

func main() {

	var number int
	fmt.Println("Enter number with 4 digits in it:")
	fmt.Scan(&number)

	fmt.Printf("Your number is: %d\n", number)

	if number < 1000 || number > 9999 {
		fmt.Println("Your number is not right")
		return
	}

	firstDigit := number / 1000
	secondDigit := (number % 1000) / 100
	thirdDigit := (number % 100) / 10
	fourthDigit := number % 10

	if firstDigit == fourthDigit && secondDigit == thirdDigit {
		fmt.Printf("%d - palindrome\n", number)
	} else {
		fmt.Printf("%d - not palindrome\n", number)
	}
}
