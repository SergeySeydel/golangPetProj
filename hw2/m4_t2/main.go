package main

import(
	"fmt"
	"sayBottles/countBottles"
)

func main(){

	var a int
	fmt.Println("Please enter quanity of bottles:")
	fmt.Scan(&a)

	if a < 0 || a > 200{
		fmt.Println("Please enter correct quanity of bottles")
		return
	}
	result := countBottles.CountBottles(a)
	fmt.Printf("In: %d\nOut: %d %s\n", a, a, result)
}