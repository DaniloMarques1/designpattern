package main

import "fmt"

type PaymentType interface {
        DoPayment(value float64)
}

type DebitPaymentType struct {
}

func (dpt DebitPaymentType) DoPayment(value float64) {
        fmt.Printf("Pagamento de %v com o cartão de débito\n", value)
}

type CreditPaymentType struct {
}

func (dpt CreditPaymentType) DoPayment(value float64) {
        fmt.Printf("Pagamento de %v com o cartão de crédito\n", value)
}
