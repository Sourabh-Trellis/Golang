package test

import (
	"Q7/app"
	"testing"
)

func TestSumArray(t *testing.T) {
	expected := 16
	actual := app.SumArray([5]int{1, 2, 3, 4, 5})

	if actual != expected {
		t.Errorf("expected %v got %v.", expected, actual)
	}
}
