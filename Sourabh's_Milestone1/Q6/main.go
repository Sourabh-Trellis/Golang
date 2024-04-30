package main

import (
	customerror "Q6/customError"
	getdayfunc "Q6/getDayFunc"
	"fmt"
)

func main() {

	var dayOfWeek int
	fmt.Println("Enter day of week (1 for Monday, 2 for Tuesday.etc)")
	fmt.Scan(&dayOfWeek)

	day, err := getdayfunc.GetDay(dayOfWeek)
	if err != nil {
		var e customerror.Error
		e.Status = false
		e.Message = err.Error()
		e.Code = 302
		e.Data = ""
		fmt.Println(e.Error())
		return
	} else {
		fmt.Println(day)
	}
}
