package my_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type MockUser struct {
	User
	mockedField string // Add a mockable field
}

func (m *MockUser) GetUsername() string {
	// Override GetUsername to return a mocked value
	return m.mockedField
}

func TestFunctionUsingUser(t *testing.T) {
	mockUser := &MockUser{User: User{ID: 1}, mockedField: "Mocked Username"}

	// Use mockUser in your function under test
	result := mockUser.GetUsername() //someFunction(mockUser) // Replace with your function

	// Assertions based on mockUser's fields or overridden methods
	assert.Equal(t, "Mocked Username", result)
}
