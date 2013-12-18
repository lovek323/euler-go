package main

import (
    "fmt"
    "./lib"
)

func isAbundant(n int) bool {
    s := 0
    for _, d := range lib.GetProperDivisors(n) {
        s += d
    }
    return s > n
}

func main() {
    var as []int
    var bs [28123]int
    s := 0

    for i := 1; i <= 28123; i++ {
        if isAbundant(i) {
            as = append(as, i)
        }
        bs[i-1] = i
        s += i
    }

    // now, find all sums of abundant numbers that are less than 28123
    N: for _, n := range as {
        M: for _, m := range as {
            if n + m > 28123 {
                // no need to continue, all subsequent values for m will be
                // greater than our upper bound
                continue N
            }
            if m < n || bs[n+m-1] == 0 {
                continue M
            }
            // remove the number from our list
            bs[n+m-1] = 0
            s -= n+m
        }
    }

    fmt.Printf("Result: %d\n", s)
}
