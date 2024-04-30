package customerror

import (
	"fmt"
)

type Error struct {
	Status  bool
	Message string
	Code    int
	Data    string
}

func (e Error) Error() string {
	return fmt.Sprintf("Status: %v\nError Message: %v\nError Code: %v\nData: %v", e.Status, e.Message, e.Code, e.Data)
}
