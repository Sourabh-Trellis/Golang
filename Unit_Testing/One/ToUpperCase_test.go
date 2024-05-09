package one

import (
	"fmt"
	"testing"
)

func TestToUpperCase(t *testing.T) {
	fmt.Println("In testing function")

	actual := ToUpperCase("sourabh")
	expected := "SOURABH"

	if actual != expected {
		t.Errorf("Expected %v ;Actual %v", expected, actual)
	}
}
