package test

import (
	"Q11/app"
	"testing"
)

func TestAreaOfCircle(t *testing.T) {

	c := app.Circle{}
	c.Radius = 5
	var expected float32 = 78.54
	actual := c.Area()

	if actual != expected {
		t.Errorf("Expected %v and got %v.", expected, actual)
	}

}

func TestAreaOfRectangle(t *testing.T)  {
	r := app.Rectangle{}
    r.Breadth = 10
	r.Length = 5
	actual := r.Area()
	var expected float32 = 50
	
	if expected != actual {
		t.Errorf("Expected %v and got %v.",expected,actual)
	}
}
