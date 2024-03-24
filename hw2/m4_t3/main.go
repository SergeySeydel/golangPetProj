package main

import(
	"fmt"
	"countTriangle/canFormTriangle"
	"countTriangle/isRightTriangle"
	"countTriangle/triangleArea"

)

func main(){

	var x1, y1, x2, y2, x3, y3 float64

	fmt.Println("Enter coordinates of free points:")

	fmt.Println("Point 1:")
	fmt.Print("x1: ")
	fmt.Scan(&x1)
	fmt.Print("y1: ")
	fmt.Scan(&y1)

	if x1 < 0 || y1 < 0{
		fmt.Println("Please enter number that equals to 0 or greater")
		return
	}

	fmt.Println("Point 2:")
	fmt.Print("x2: ")
	fmt.Scan(&x2)
	fmt.Print("y2: ")
	fmt.Scan(&y2)

	if x2 < 0 || y2 < 0{
		fmt.Println("Please enter number that equals to 0 or greater")
		return
	}

	fmt.Println("Point 3:")
	fmt.Print("x3: ")
	fmt.Scan(&x3)
	fmt.Print("y3: ")
	fmt.Scan(&y3)

	if x3 < 0 || y3 < 0{
		fmt.Println("Please enter number that equals to 0 or greater")
		return
	}

	
	if canFormTriangle.CanFormTriangle(x1, y1, x2, y2, x3, y3){
		fmt.Println("Triangle can be formed")
		fmt.Println("Area of triangle:", triangleArea.TriangleArea(x1, y1, x2, y2, x3, y3))
		fmt.Println("Is the triangle right-angled?", isRightTriangle.IsRightTriangle(x1, y1, x2, y2, x3, y3))
	}else{
		fmt.Println("Triangle can't be formed")
		return
	}
	
}