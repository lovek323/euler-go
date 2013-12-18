package main

import (
    "fmt"
    "math"
)

const n int64 = 600851475143

func isPrime(x int64) bool {
    var i int64

    for i = 2; i <= int64(math.Sqrt(float64(x))); i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}

func main() {
    var factor int64
    var i int64

    for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            // we have found a factor, is it prime
            if isPrime(i) {
                factor = i
            }
        }
    }

    fmt.Printf("Largest factor: %d\n", factor)
}
