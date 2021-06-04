// context class
package main

type Calculator struct {
        strategy Arithmetic
}

func NewCalculator(strategy Arithmetic) *Calculator {
        return &Calculator{
                strategy: strategy,
        }
}

func (c *Calculator) ChangeStrategy(strategy Arithmetic)  {
        c.strategy = strategy
}

func (c *Calculator) Execute(a, b float64) float64 {
        return c.strategy.Execute(a, b)
}
