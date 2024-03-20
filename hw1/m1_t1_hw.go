package main

import "fmt"

func main() {
	const costPerLiter = 48 

	var distance, consumption float64

	fmt.Println("Введите расстояние в километрах:")
	fmt.Scanln(&distance)

	fmt.Println("Введите расход топлива в литрах на 100 км:")
	fmt.Scanln(&consumption)

	if distance < 50 || distance > 10000 || consumption < 5 || consumption > 25 {
		fmt.Println("Пожалуйста, введите корректные данные.")
		return
	}

	totalCost := (distance / 100) * consumption * costPerLiter
	fmt.Printf("Стоимость поездки в рублях: %.2f", totalCost)
}
