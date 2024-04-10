package main
 
import "fmt"
 
func main() {
    var number int
    fmt.Println("Enter a number that you want to check for fibonacci")
    fmt.Scan(&number)
    a := 0
    b := 1
    //var c int
    for i := 0; i <= number; i++ {
        fmt.Println(a)
        a, b = b, a+b
    }
}
 