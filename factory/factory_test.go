package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid with cash") {
		t.Error("Paid with error")
	}

	t.Log("Log", msg)
}

func TestCreatePaymentMethodDebit(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatal("")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid with debit card") {
		t.Error("")
	}

	t.Log("Log", msg)
}
