package test

import (
	getdayfunc "Q6/getDayFunc"
	"testing"
)

func TestGetDay(t *testing.T) {
	actual,_ := getdayfunc.GetDay(3)
	expected := "Day is monday"
	if expected != actual {
		t.Errorf("Epected '%v' but got '%v'.",expected,actual)
	}
}
