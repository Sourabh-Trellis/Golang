package main

import "fmt"

type product struct {
	productId   int
	productname string
	price       int
}

func main() {
	fmt.Print("\033[H\033[2J")
	var productList = make(map[int]product)

	var count int = 100

	for {
		fmt.Println("-----Welcome to inventory management system-----")
		fmt.Println("\t1)Add product.\n\t2)List products.\n\t3)Delete product.\n\t4)Update product details.\n\t5)Exit")

		var choice int
		fmt.Print("Choose any option:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			p := new(product)

			p.productId = count

			fmt.Println("Enter product Name:")
			fmt.Scan(&p.productname)

			fmt.Println("Enter product price:")
			fmt.Scan(&p.price)

			productList[p.productId] = *p
			count++
		case 2:
			for k, v := range productList {
				fmt.Printf("%v : %v\n", k, v)
			}
		case 3:
			var id int
			fmt.Println("Enter productId to delete product:")
			fmt.Scan(&id)

			if value, ok := productList[id]; ok {
				delete(productList, id)
				fmt.Printf("Product \"%v\" deleted\n", value.productname)
			} else {
				fmt.Println("Product was not found.")
			}

		case 4:
			var name string
			fmt.Println("Enter product name to update :")
			fmt.Scan(&name)
			for productId, product := range productList {
				if product.productname == name {
					fmt.Println("Enter updated price:")
					fmt.Scan(&product.price)
					productList[productId] = product
					fmt.Println("Updated successfully")
				} else {
					fmt.Println("Product not found.")
				}
			}
		case 5:
			return
		}
	}

}
