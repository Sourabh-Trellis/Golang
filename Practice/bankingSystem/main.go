package main

import "fmt"

func main() {
	var accNum int32 = 174832
	bank := make(map[int32]account)
	fmt.Println("Welcome to bank")
	fmt.Println("Choose from following options:")
	for {
		fmt.Println("------------------------------")
		var choice int
		fmt.Println("1)Open new Account\n2)Withdraw Amount\n3)Deposit amount\n4)Check balance\n5)Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			acc := openAccount(accNum)
			accNum++
			bank[acc.getAccNumber()] = acc

		case 2:
			var acc int32
			fmt.Println("Enter your accout number")
			fmt.Scan(&acc)
			if _,ok := bank[acc]; ok {
				withdraw(&bank,acc)
			} else {
				fmt.Println("Invalid Account Number.")
			}
			
		case 3:
		case 4:
		case 5:

		}
	}
}
