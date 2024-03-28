package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Invoice struct {
	ProductName  string
	Quantity     int
	BuyerName    string
	ContactPhone string
	Address      Address
}

type Address struct {
	Index  string
	City   string
	Street string
	Number string
	Flat   string
}

func (inv *Invoice) Validate() string {
	if inv.isEmpty() {
		return "все поля должны быть заполнены"
	}
	if !isAlpha(inv.BuyerName) {
		return "ФИО покупателя должно содержать только буквы"
	}
	if len(inv.ContactPhone) != 10 || !isNumeric(inv.ContactPhone) {
		return "некорректный формат телефонного номера (должно быть 10 цифр)"
	}
	if len(inv.Address.Index) != 6 || !isNumeric(inv.Address.Index) {
		return "некорректный формат индекса (должно быть ровно 6 цифр)"
	}
	if !isAlpha(inv.Address.City) {
		return "Город должен содержать только буквы"
	}
	if !isAlpha(inv.Address.Street) {
		return "Улица должна содержать только буквы"
	}
	if !isNumeric(inv.Address.Number) {
		return "Дом должен содержать только цифры"
	}
	if !isNumeric(inv.Address.Flat) {
		return "Квартира должна содержать только цифры"
	}
	return ""
}

func (inv *Invoice) isEmpty() bool {
	return inv.ProductName == "" || inv.Quantity == 0 || inv.BuyerName == "" || inv.ContactPhone == "" ||
		inv.Address.Index == "" || inv.Address.City == "" || inv.Address.Street == "" ||
		inv.Address.Number == "" || inv.Address.Flat == ""
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isAlpha(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func (inv *Invoice) Print() {
	fmt.Println("Накладная:")
	fmt.Println("Наименование товара:", inv.ProductName)
	fmt.Println("Количество:", inv.Quantity)
	fmt.Println("ФИО покупателя:", inv.BuyerName)
	fmt.Println("Контактный телефон:", inv.ContactPhone)
	fmt.Println("Адрес:", inv.Address.Index, inv.Address.City, inv.Address.Street, inv.Address.Number, inv.Address.Flat)
}

func NewInvoice(productName string, quantity int, buyerName string, contactPhone string, address Address) *Invoice {
	return &Invoice{
		ProductName:  productName,
		Quantity:     quantity,
		BuyerName:    buyerName,
		ContactPhone: contactPhone,
		Address:      address,
	}
}

func main() {
	var productName, buyerName, contactPhone, index, city, street, number, flat string
	var quantity int

	fmt.Println("Введите информацию о накладной:")

	fmt.Println("Наименование товара (минимум 1, максимум 100 символов):")
	fmt.Scanln(&productName)
	for len(productName) < 1 || len(productName) > 100 {
		fmt.Println("Ошибка: Наименование товара должно содержать от 1 до 100 символов. Введите снова:")
		fmt.Scanln(&productName)
	}

	fmt.Println("Количество:")
	for {
		_, err := fmt.Scanln(&quantity)
		if err == nil && quantity > 0 {
			break
		}
		fmt.Println("Ошибка: Количество должно быть положительным числом. Введите снова:")
	}

	fmt.Println("ФИО покупателя (только буквы):")
	fmt.Scanln(&buyerName)
	for !isAlpha(buyerName) {
		fmt.Println("Ошибка: ФИО покупателя должно содержать только буквы. Введите снова:")
		fmt.Scanln(&buyerName)
	}

	fmt.Println("Контактный телефон (10 цифр):")
	fmt.Scanln(&contactPhone)
	for len(contactPhone) != 10 || !isNumeric(contactPhone) {
		fmt.Println("Ошибка: Некорректный формат телефонного номера. Введите снова:")
		fmt.Scanln(&contactPhone)
	}

	fmt.Println("Адрес:")
	fmt.Println("Индекс (ровно 6 цифр):")
	fmt.Scanln(&index)
	for len(index) != 6 || !isNumeric(index) {
		fmt.Println("Ошибка: Некорректный формат индекса. Введите снова:")
		fmt.Scanln(&index)
	}

	fmt.Println("Город (только буквы):")
	fmt.Scanln(&city)
	for !isAlpha(city) {
		fmt.Println("Ошибка: Город должен содержать только буквы. Введите снова:")
		fmt.Scanln(&city)
	}

	fmt.Println("Улица (только буквы):")
	fmt.Scanln(&street)
	for !isAlpha(street) {
		fmt.Println("Ошибка: Улица должна содержать только буквы. Введите снова:")
		fmt.Scanln(&street)
	}

	fmt.Println("Дом (только цифры):")
	fmt.Scanln(&number)
	for !isNumeric(number) {
		fmt.Println("Ошибка: Дом должен содержать только цифры. Введите снова:")
		fmt.Scanln(&number)
	}

	fmt.Println("Квартира (только цифры):")
	fmt.Scanln(&flat)
	for !isNumeric(flat) {
		fmt.Println("Ошибка: Квартира должна содержать только цифры. Введите снова:")
		fmt.Scanln(&flat)
	}

	address := Address{
		Index:  index,
		City:   city,
		Street: street,
		Number: number,
		Flat:   flat,
	}

	inv := NewInvoice(productName, quantity, buyerName, contactPhone, address)

	errMsg := inv.Validate()
	if errMsg != "" {
		fmt.Println("Ошибка валидации накладной:", errMsg)
		return
	}

	inv.Print()
}
