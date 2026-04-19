package main

import (
	"fmt"
)

type ProductID int
type Price int
type Months int

type Product struct {
	ID     ProductID
	Name   string
	Price  Price
	Months Months
}

func monthlyPayment(p Product) int {
	if p.Months < 3 || p.Months > 24 {
		return 0
	}
	return int(p.Price) / int(p.Months)
}

func printProduct(p Product) {
	fmt.Println("===== Alif Shop =====")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Товар: %s\n", p.Name)
	fmt.Printf("Цена: %d сум\n", p.Price)
	fmt.Printf("Рассрочка: %d месяцев\n", p.Months)
	fmt.Printf("В месяц: %d сум\n", monthlyPayment(p))
	fmt.Println("=====================================")
}

func main() {
	product := Product{
		ID:     7291,
		Name:   "Xiaomi Redmi Note 14",
		Price:  2499000,
		Months: 12,
	}

	printProduct(product)
}
