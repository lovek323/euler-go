package main

import (
    "fmt"
)

func sumSquares(n int) int {
    sum := 0

    for i := 1; i <= n; i++ {
        sum += i * i
    }

    return sum
}

func squareSum(n int) int {
    sum := 0

    for i := 1; i <= n; i++ {
        sum += i
    }

    return sum * sum
}

func main() {
    fmt.Printf("Difference: %d\n", sumSquares(100) - squareSum(100))
}
