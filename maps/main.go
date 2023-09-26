package main

import (
	"fmt"
)

type Product struct {
	Title  string
	Price int
	Stock int
}



func main() {

	productStore := make (map[int]Product)
	nextID := 1

	for {
		fmt.Println("1. Add Product")
		fmt.Println("2. Delete Product")
		fmt.Println("3. Edit Product")
		fmt.Println("4. Print Products")
		fmt.Println("5. Exit")

		var choice int;
		fmt.Println("Enter your Choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title string
			var price, stock int
			fmt.Print("Enter title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter stock: ")
			fmt.Scanln(&stock)

			productStore[nextID] = Product{
				Title: title,
				Price: price,
				Stock: stock,
			}
			nextID++

		case 2:
			var id int
			fmt.Print("Enter product ID to be deleted: ")	
			fmt.Scanln(&id)

			if _, ok := productStore[id], ok {
				delete(productStore, id)
				fmt.Println("product deleted.")
			} else {
				fmt.Println("Product not found.")
			}
		}

	}
}