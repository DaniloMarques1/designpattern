package main

import "fmt"
func main() {
        somar := AddStrategy{}
        calculator := NewCalculator(somar)
        fmt.Println(calculator.Execute(2, 3))

        sub := SubStrategy{}
        calculator.ChangeStrategy(sub)
        fmt.Println(calculator.Execute(2, 3))
}
