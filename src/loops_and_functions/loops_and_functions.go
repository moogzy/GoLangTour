package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 10.0
    for v := 0 ; v < 10 ; v++ {
        z = z - ((math.Pow(z,2) - x) / (2 * z))
    }
    return z
}

func main() {
    value := 0.0
    for count := 0 ; count < 10 ; count++ {
        fmt.Printf("My function(%d): %g\n", int(value), Sqrt(value))
        fmt.Printf("Math Sqrt package(%d): %g\n", int(value), math.Sqrt(value))
        value++
    }
}
