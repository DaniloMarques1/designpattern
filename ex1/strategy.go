// startegies
package main

type Arithmetic interface {
        Execute(a, b float64) float64
}

type AddStrategy struct {
}

func (add AddStrategy) Execute(a, b float64) float64 {
        return a + b
}

type SubStrategy struct {
}

func (sub SubStrategy) Execute(a, b float64) float64 {
        return a - b
}
