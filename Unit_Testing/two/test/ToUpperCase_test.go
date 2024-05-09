package test

import (
	"fmt"
	one "test2/app"
	"testing"
)

func TestToUpperCase(t *testing.T) {
	fmt.Println("In testing function")

	actual := one.ToUpperCase("sourabh")
	expected := "SOURABH"

	if actual != expected {
		t.Errorf("Expected %v ;Actual %v", expected, actual)
	}
}
