package main

import "fmt"

// product management system

type Products struct {
	Name     string
	Price    int
	Category string
	InStock  bool
}

func printProducts(Product []Products) {
	for _, Products := range Product {
		fmt.Println("name", Products.Name)
		fmt.Println("price", Products.Price)
		fmt.Println("category", Products.Category)
		fmt.Println("instock", Products.InStock)

		switch Products.Category {
		case "E":
			fmt.Println("electronics ")
		case "C":
			fmt.Println("clothing")
		case "F":
			fmt.Println("foods")
		default:
			fmt.Println("other category ")
		}
		fmt.Println()
	}

}

func main() {
	Product := []Products{
		{Name: "Milk", Price: 62, Category: "F", InStock: true},
		{Name: "t-shirt", Price: 499, Category: "C", InStock: true},
		{Name: "Phone", Price: 15000, Category: "E", InStock: false},
		{Name: "Macbook", Price: 149000, Category: "E", InStock: true},
	}

	printProducts(Product)

	stock := map[string]int{
		"Instock":    0,
		"OutOfStock": 0,
	}
	for _, Products := range Product {
		if Products.InStock {
			stock["Instock"]++
		} else {
			stock["OutOfStock"]++
		}
	}

	fmt.Println("Stock Summary")
	fmt.Println("in_stock", stock["Instock"])
	fmt.Println("out_of_stock", stock["OutOfStock"])
}
