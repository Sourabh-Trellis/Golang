package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var accNum int32 = 174832
	bank := make(map[int32]account)
	fmt.Println("Welcome to bank")
	fmt.Println("Choose from following options:")
	for {
		fmt.Println("------------------------------")
		var choice int
		fmt.Println("1)Open new Account\n2)Withdraw Amount\n3)Deposit amount\n4)Check balance\n5)Show all accounts\n6)Exit")
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
			if _, ok := bank[acc]; ok {
				withdraw(&bank, acc)
			} else {
				fmt.Println("Invalid Account Number.")
			}

		case 3:
		case 4:
		case 5:
		case 6:

		}
	}
}

// ////////////////////////////////////////////////////////
type account struct {
	accNumber int32
	name      string
	acctype   string
	balance   float64
}

func (account *account) setAccNumber(n int32) {
	account.accNumber = n
}

func (account *account) setName(name string) {
	account.name = name
}

func (account *account) setAccType(Type string) {
	account.acctype = Type
}

func (account *account) setBalance(balance float64) {
	account.balance = balance
}

func (account *account) getAccNumber() int32 {
	return account.accNumber
}

func openAccount(accNum int32) account {
	reader := bufio.NewReader(os.Stdin)
	acc := account{}

	acc.setAccNumber(accNum)
	fmt.Println("Enter your name:")
	var name string
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	acc.setName(name)

	fmt.Println("enter account type from the following ( Current / Savings ):")
	Type, _ := reader.ReadString('\n')
	Type = strings.TrimSpace(Type)
	acc.setAccType(Type)

	acc.setBalance(1000)
	return acc
}

func withdraw(bank *map[int32]account, accnumber int32) {
	var amount float64
	fmt.Println("Enter ammount to withdraw :")
	fmt.Scan(&amount)
	holder := (*bank)[accnumber]
	if holder.balance-amount > 1000 {
		holder.balance = holder.balance - amount
	} else {
		fmt.Println("Insufficient Balance.Account should maintain minimum balance Rs.1000")
	}
	fmt.Println("Current balance is", holder.balance)
}

func deposit(bank *map[int32]account, accnumber int32) {
	var amount float64
	fmt.Println("Enter ammount to deposit :")
	fmt.Scan(&amount)
	holder := (*bank)[accnumber]
	holder.balance = holder.balance + amount
	fmt.Println("Amount deposited successfully.")
	fmt.Println("Current balance is", holder.balance)
}
