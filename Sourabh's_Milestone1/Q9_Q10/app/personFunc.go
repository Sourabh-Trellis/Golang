package app

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) ShowDetails() string {
	// fmt.Println("Name :", p.Name,"Age :", p.Age)
	return "Name :"+ p.Name+" Age :"+ fmt.Sprintf("%d", p.Age)
	// fmt.Println("Age :", p.age)
}

func (p Person) IsAdult() bool {
	if p.Age >= 18 {
		return true
	} else {
		return false
	}
}
