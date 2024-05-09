package test

import (
	app "Q3/app"
	"testing"
)

//8,2,1,15
func TestBasicMath(t *testing.T)  {
	var expectedSum,expectedSub,expectedDiv,expectedMul float32 = 8,2,1,15 
	actualSum,actualSub,actualDiv,actualMul := app.BasicMath(5,3)
	if actualSum != expectedSum || expectedSub != actualSub || expectedDiv != actualDiv || expectedMul != actualMul{
		t.Errorf("Expected values are %v,%v,%v,%v and got %v,%v,%v,%v.",expectedSum,expectedSub,expectedDiv,expectedMul,actualSum,actualSub,actualDiv,actualMul)
	}
}