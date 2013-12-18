package main

import (
    "fmt"
    "./lib"
)

func getSumProperDivisors(n int) int {
    d := lib.GetProperDivisors(n)
    s := 0

    for _, di := range d {
        s += di
    }

    return s
}

func isAmicable(n int) bool {
    sn := getSumProperDivisors(n)

    if sn == n {
        return false
    }

    sm := getSumProperDivisors(sn)

    return sm == n
}

func main() {
    s := 0

    for i := 1; i < 10000; i++ {
        if isAmicable(i) {
            s += i
        }
    }

    fmt.Printf("Result: %d\n", s)
}
