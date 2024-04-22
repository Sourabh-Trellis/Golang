package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) showDetails() {
	fmt.Println("Name :", p.name)
	fmt.Println("Age :", p.age)
}

func (p Person) isAdult() bool {
	if p.age >= 18 {
		return true
	} else {
		return false
	}
}

func main() {

	person := Person{"Sourabh", 24}
	person.showDetails()

	fmt.Println("Is Adult :",person.isAdult())
}
