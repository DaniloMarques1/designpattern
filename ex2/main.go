package main

func main() {
        creditPayment := CreditPaymentType{}
        paymentModule := PaymentModule{paymentType: creditPayment}
        paymentModule.DoPayment(30)

        debitPayment := DebitPaymentType{}
        paymentModule.ChangeStrategy(debitPayment)
        paymentModule.DoPayment(30)
}
