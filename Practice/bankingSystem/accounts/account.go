package accounts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Account struct {
	accNumber int32
	name      string
	acctype   string
	balance   float64
}

func (account *Account) SetAccNumber(n int32) {
	account.accNumber = n
}

func (account *Account) SetName(name string) {
	account.name = name
}

func (account *Account) SetAccType(Type string) {
	account.acctype = Type
}

func (account *Account) SetBalance(balance float64) {
	account.balance = balance
}

func (account *Account) GetName() string {
	return account.name
}

func (account *Account) GetAccNumber() int32 {
	return account.accNumber
}

func (account *Account) GetAccType() string {
	return account.acctype
}

func (account *Account) GetBalance() float64 {
	return account.balance
}

func OpenAccount(accNum int32) Account {
	reader := bufio.NewReader(os.Stdin)
	acc := Account{}

	acc.SetAccNumber(accNum)
	fmt.Println("Enter your name:")
	var name string
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	acc.SetName(name)

	fmt.Println("enter account type from the following ( Current / Savings ):")
	Type, _ := reader.ReadString('\n')
	Type = strings.TrimSpace(Type)
	acc.SetAccType(Type)

	acc.SetBalance(1000)
	return acc
}

func Withdraw(bank *map[int32]Account, accnumber int32) {
	var amount float64
	fmt.Println("Enter ammount to withdraw :")
	fmt.Scan(&amount)
	holder := (*bank)[accnumber]
	holder.balance = holder.balance - amount
	(*bank)[accnumber] = holder	
	fmt.Println("Current balance is", holder.balance)
}

func Deposit(bank *map[int32]Account, accnumber int32) {
	var amount float64
	fmt.Println("Enter ammount to deposit :")
	fmt.Scan(&amount)
	holder := (*bank)[accnumber]
	holder.balance = holder.balance + amount
	(*bank)[accnumber] = holder
	fmt.Println("Amount deposited successfully.")
	fmt.Println("Current balance is", holder.balance)
}

func CheckBalance(bank *map[int32]Account, accaccount int32) {
	holder := (*bank)[accaccount]
	fmt.Printf("Current balance is %v.", holder.balance)
}
