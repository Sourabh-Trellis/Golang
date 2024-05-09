package test

import (
	app "Q2/app"
	"testing"
)

func TestPrintPI(t *testing.T) {
	expected := 3.14
	actual := app.PrintPI()
	if actual != expected {
		t.Errorf("Pi should be %v got %v.", expected, actual)
	}
}
