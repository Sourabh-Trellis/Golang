package test

import (
	"Q9/app"
	"testing"
)

func TestIsAdult(t *testing.T) {
	person := app.Person{}
	person.Name = "Sourabh"
	person.Age = 17
	actual := person.IsAdult()
	expected := true

	if actual != expected {
		t.Errorf("Expected %v got %v.",expected,actual)
	} 
}

func TestShowDetails(t *testing.T)  {
	person := app.Person{}
	person.Name = "Sourabh"
	person.Age = 24
	expected := "Name :Sourabh Age :2"
	actual := person.ShowDetails()

	if actual != expected {
		t.Errorf("Expected %v and got %v.",expected,actual)
	}
}
