package errors

import "fmt"

type DivideByZeroError struct {
	ErrMessage string
	Code       int
}

func (e *DivideByZeroError) Error() string {
	return fmt.Sprintf("Error message :%v\nError Code :%v",e.ErrMessage,e.Code)
}
