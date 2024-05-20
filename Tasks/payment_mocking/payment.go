package paymentmocking

type PaymentGateway interface {
	ProcessPayment(amount float64, card_number int, exp_date string, CVV int) (bool, error)
}

type PaymentService struct {
	Service PaymentGateway
}

func NewPaymentService(gateway PaymentGateway) *PaymentService {
	return &PaymentService{Service: gateway}
}

func (p *PaymentService) MakePayment(amount float64, card_number int, exp_date string, CVV int) (bool, error) {
	return p.Service.ProcessPayment(amount, card_number, exp_date, CVV)
}
