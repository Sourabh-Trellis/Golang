package test

import (
	print "Q1/app"
	"testing"
)

func TestPrintName(t *testing.T) {
	name := print.PrintName("john")
	expected := "john"
	if name != expected {
		t.Errorf("Name should be %v got %v", expected, name)
	}
}
