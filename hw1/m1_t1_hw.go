package main

import "fmt"

/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/


func main() {
	const costPerLiter = 48 

	var distance, consumption float64

	fmt.Println("Enter distance in kilometers:")
	fmt.Scanln(&distance)

	fmt.Println("Enter fuel consumption for each 100 kilometers:")
	fmt.Scanln(&consumption)

	if distance < 50 || distance > 10000 || consumption < 5 || consumption > 25 {
		fmt.Println("Please enter correct information")
		return
	}

	totalCost := (distance / 100) * consumption * costPerLiter
	fmt.Printf("Cost of trail is: %.2f rubles\n", totalCost)
}
