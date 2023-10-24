package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitPM), nil
	default:
		return nil, errors.New("Error")
	}
}

type CashPM struct{}
type DebitPM struct{}

func (p *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid with cash\n", amount)
}

func (p *DebitPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid with debit card\n", amount)
}
