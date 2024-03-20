package main

import "fmt"

func main(){

	var a, b, c int

	fmt.Println("Enter 3 digits separated by space key:")
	fmt.Scan(&a, &b, &c)

	fmt.Printf("Digits input: %d %d %d\n", a, b, c)

	if a > b{
		a, b = b, a 
	}
	if a > c{
		a, c = c, a
	}
	if b > c{
		b, c = c, b
	}
	
	fmt.Printf("Digits output in ascending order: %d %d %d\n", a, b, c)
}