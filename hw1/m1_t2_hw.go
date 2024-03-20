package main

import "fmt"

func main() {
  
  var number, lastDigit, middleDigit, firstdigit int

  fmt.Println("Enter 3 digit number:")
  fmt.Scan(&number)

  lastDigit = number % 10
  middleDigit = (number / 10) % 10
  firstDigit = number / 100

  fmt.Printf("Input number is: %d, Output number is: %d%d%d\n", number, lastDigit, middleDigit, firstDigit)
}
