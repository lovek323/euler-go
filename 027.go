package main

import (
    "fmt"
    "./lib"
)

func getPrimeCount(a, b int) int {
    pc := 0
    poly := func (n int) int {
        return n*n+a*n+b
    }
    for n := 0; ; n++ {
        pn := poly(n)
        if !lib.IsPrime(pn) {
            break
        }
        pc++
    }
    return pc
}

func main() {
    lpc := 0
    lm := 0

    for a := -1001; a < 1000; a++ {
        if a % 2 == 0 {
            // a must be odd
            continue
        }
        for b := -1001; b < 1000; b++ {
            if !lib.IsPrime(b) {
                // b must be prime, so we can skip
                continue
            }
            pc := getPrimeCount(a, b)
            if pc > lpc {
                lpc = pc
                lm = a*b
            }
        }
    }

    fmt.Printf("Result: %d\n", lm)
}
