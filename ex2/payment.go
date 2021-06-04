package main

type PaymentModule struct {
        paymentType PaymentType
}

func (pm *PaymentModule) ChangeStrategy(paymentType PaymentType) {
        pm.paymentType = paymentType
}

func (pm *PaymentModule) DoPayment(value float64) {
        pm.paymentType.DoPayment(value)
}
