package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type account struct {
// 	accNumber int32
// 	name      string
// 	acctype   string
// 	balance   float64
// }

// func (account *account) setAccNumber(n int32) {
// 	account.accNumber = n
// }

// func (account *account) setName(name string) {
// 	account.name = name
// }

// func (account *account) setAccType(Type string) {
// 	account.acctype = Type
// }

// func (account *account) setBalance(balance float64) {
// 	account.balance = balance
// }

// func (account *account) getAccNumber() int32 {
// 	return account.accNumber
// }

// func openAccount(accNum int32) account {
// 	reader := bufio.NewReader(os.Stdin)
// 	acc := account{}

// 	acc.setAccNumber(accNum)
// 	fmt.Println("Enter your name:")
// 	var name string
// 	name, _ = reader.ReadString('\n')
// 	name = strings.TrimSpace(name)
// 	acc.setName(name)

// 	fmt.Println("enter account type from the following ( Current / Savings ):")
// 	Type, _ := reader.ReadString('\n')
// 	Type = strings.TrimSpace(Type)
// 	acc.setAccType(Type)

// 	acc.setBalance(1000)
// 	return acc
// }

// func withdraw(bank *map[int32]account, accnumber int32) {
// 	var amount float64
// 	fmt.Println("Enter ammount to withdraw :")
// 	fmt.Scan(&amount)
// 	holder := (*bank)[accnumber]
// 	holder.balance = holder.balance - amount
// 	fmt.Println("Current balance is", holder.balance)
// }
