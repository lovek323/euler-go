package main

import (
    "fmt"
    "math"
)

var primes []int

func nextPrime() int {
    last := primes[len(primes)-1]

    // fmt.Printf("Starting at %d\n", last)

    N: for i := last+2; ; i += 2 {
        for _, p := range primes {
            if i % p == 0 {
                continue N
            }
            if p > int(math.Sqrt(float64(i))) {
                break
            }
        }
        primes = append(primes, i)
        break
    }

    return primes[len(primes)-1]
}

func main() {
    primes = append(primes, 2, 3)
    prime := 0

    for ; prime < 2e6 ; prime = nextPrime() { }

    sum := 0

    for i := 0; i < len(primes)-1; i++ {
        sum += primes[i]
    }

    fmt.Printf("Sum: %d\n", sum)
}
