package main

import (
    "fmt"
    "./lib"
)

const n int64 = 600851475143

func main() {
    var factor int64
    var i int64

    for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            // we have found a factor, is it prime
            if lib.IsPrime64(i) {
                factor = i
            }
        }
    }

    fmt.Printf("Largest factor: %d\n", factor)
}
