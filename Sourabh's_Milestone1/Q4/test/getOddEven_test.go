package test

import "testing"
import app "Q4/app"

func TestGetOddEven(t *testing.T)  {
	expected := true
	actual := app.GetOddEven(3)
	
	if actual != expected {
		t.Errorf("Expected value is %v but got %v.",expected,actual)
	}
}