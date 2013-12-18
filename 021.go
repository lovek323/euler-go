package main

import "fmt"

func getProperDivisors(n int) []int {
    var d []int

    for i := 1; i < n; i++ {
        if n % i == 0 {
            d = append(d, i)
        }
    }

    return d
}

func getSumProperDivisors(n int) int {
    d := getProperDivisors(n)
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
