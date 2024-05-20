package paymentmocking

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPaymentGateway struct {
	mock.Mock
}

func (m *mockPaymentGateway) ProcessPayment(amount float64, card_number int, exp_date string, CVV int) (bool, error) {
	args := m.Called(amount, card_number, exp_date, CVV)
	return args.Bool(0), args.Error(1)
}

func TestPayment_success(t *testing.T) {
	mockPay := new(mockPaymentGateway)

	mockPay.On("ProcessPayment", 2834.50, 1846382624196946, "17-05-2024", 258).Return(true, nil)

	PayService := NewPaymentService(mockPay)

	status, err := PayService.MakePayment(2834.50, 1846382624196946, "17-05-2024", 258)

	assert.Equal(t, true, status)
	assert.NoError(t, err)

	mockPay.AssertCalled(t, "ProcessPayment", 2834.50, 1846382624196946, "17-05-2024", 258)
}

func TestPayment_fail(t *testing.T) {
	mockPay := new(mockPaymentGateway)

	mockPay.On("ProcessPayment", 2736.54, 745938530215724, "15-06-2024", 165).Return(false, errors.New("Invalid Card number."))
	mockPay.On("ProcessPayment", 7456.36, 1637295623406590, "14-03-2023", 850).Return(false, errors.New("Card expired."))
	mockPay.On("ProcessPayment", 7456.36, 1637295623406590, "26-08-2024", 85054).Return(false, errors.New("Invalid CVV."))
	mockPay.On("ProcessPayment", -4646.00, 1637295623406590, "26-08-2024", 850).Return(false, errors.New("Invalid Amount."))

	PayService := NewPaymentService(mockPay)

	status, err := PayService.MakePayment(2736.54, 745938530215724, "15-06-2024", 165)
	assert.Equal(t, false, status)
	assert.Error(t, err)

	status, err = PayService.MakePayment(7456.36, 1637295623406590, "14-03-2023", 850)
	assert.Equal(t, false, status)
	assert.Error(t, err)

	status, err = PayService.MakePayment(7456.36, 1637295623406590, "26-08-2024", 85054)
	assert.Equal(t, false, status)
	assert.Error(t, err)

	status, err = PayService.MakePayment(-4646.00, 1637295623406590, "26-08-2024", 850)
	assert.Equal(t, false, status)
	assert.Error(t, err)
}
