package fgfg

import (
	"fmt"
	"testing"
	one "three/app"
)

func TestToUpperCase(t *testing.T) {
	fmt.Println("In testing function")

	actual := one.ToUpperCase("sourabh8")
	expected := "SOURABH8"
	fmt.Println(actual)
	if actual != expected {
		t.Errorf("Expected %v ;Actual %v", expected, actual)
	}
}


