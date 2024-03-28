package main

import (
	"fmt"
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
	if !isNumeric(inv.ContactPhone) {
		return "некорректный формат телефонного номера"
	}
	if !isNumeric(inv.Address.Index) {
		return "некорректный формат индекса"
	}
	return ""
}

func (inv *Invoice) isEmpty() bool {
	return inv.ProductName == "" || inv.Quantity == 0 || inv.BuyerName == "" || inv.ContactPhone == "" ||
		inv.Address.Index == "" || inv.Address.City == "" || inv.Address.Street == "" ||
		inv.Address.Number == "" || inv.Address.Flat == ""
}

func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
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

	fmt.Println("Наименование товара: ")
	fmt.Scanln(&productName)

	fmt.Println("Количество: ")
	fmt.Scanln(&quantity)

	fmt.Println("ФИО покупателя: ")
	fmt.Scanln(&buyerName)

	fmt.Println("Контактный телефон: ")
	fmt.Scanln(&contactPhone)

	fmt.Println("Адрес:")
	fmt.Println("Индекс: ")
	fmt.Scanln(&index)

	fmt.Println("Город: ")
	fmt.Scanln(&city)

	fmt.Println("Улица: ")
	fmt.Scanln(&street)

	fmt.Println("Дом: ")
	fmt.Scanln(&number)

	fmt.Println("Квартира: ")
	fmt.Scanln(&flat)

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
