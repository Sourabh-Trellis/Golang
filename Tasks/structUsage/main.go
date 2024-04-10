package main

import "fmt"

// Task 4: Struct Usage - Define a struct named Employee with fields Name, Age, and Salary. Write a program that creates an array of Employee instances and prints the details of the employee with the highest salary.

type Employee struct {
	name   string
	age    int
	salary float64
}

func main() {

	company := [5]Employee{{"sourabh", 24, 40000}, {"sangam", 24, 50000}, {"kirti", 24, 35000}, {"sanika", 24, 45000}, {"apeksha", 24, 40000}}

	highest := empWithHighestSal(company)
	fmt.Printf("Defails of employee with highest salary :\n Name \t\t Age \t\t Salary \n %v \t %v \t\t %v\n",highest.name,highest.age,highest.salary)

}

func empWithHighestSal(emps [5]Employee) Employee {

	var highestSalaried Employee

	for _, emp := range emps {
		if emp.salary > highestSalaried.salary {
			highestSalaried = emp
		}
	}
	return highestSalaried
}
