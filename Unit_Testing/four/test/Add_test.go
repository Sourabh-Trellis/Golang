package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	structs := []struct {
		a   float32
		b   float32
		sum float32
	}{
		{2, 4, 6},
		{1, 7, 8},
		{3, 5, 8.0},
	}

	for _, st := range structs {
		t.Run()
		actual := st.a + st.b
		assert.Equal(t, st.sum, actual, "Should be equal")
	}
}
