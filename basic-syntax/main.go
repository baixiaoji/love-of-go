package bookstore

import "fmt"

type Customer struct {
	Name  string
	Email string
}

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	var title string
	var copies int
	var editionVerion string
	var inStock bool
	var isSpecialOffer bool
	var royaltyPercentage float64
	var discountPercentage float64
	fmt.Println(editionVerion)
	fmt.Println(copies)
	fmt.Println(title)
	fmt.Println(isSpecialOffer)
	fmt.Println(inStock)
	fmt.Println(royaltyPercentage)
	fmt.Println(discountPercentage)
	title = "Hello, World!"
	copies = 99
	editionVerion = "1.0.0"
	inStock = true
	royaltyPercentage = 12.5
	isSpecialOffer = false
	discountPercentage = 0.7

	fmt.Println(title)
	fmt.Println(copies)
}
