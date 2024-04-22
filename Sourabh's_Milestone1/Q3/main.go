package main

import "fmt"

func basicMath(num1 ,num2 float32) (add ,sub ,div ,mul float32) {
	add = num1 + num2
	sub = num1 - num2
	div = num1 / num2
	mul = num1 * num2
	return
}

func main()  {
	
 var num1 float32
 var num2 float32
 fmt.Println("Enter Number 1 :")
 fmt.Scan(&num1)
 fmt.Println("Enter number 2 :")
 fmt.Scan(&num2)

 addition ,Substraction ,division ,Multiplication := basicMath(num1,num2)

 fmt.Println("Addition is",addition)
 fmt.Println("Substraction is",Substraction)
 fmt.Println("Division is",division)
 fmt.Println("Multiplication is",Multiplication)

}