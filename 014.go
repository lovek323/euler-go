package main

import (
    "fmt"
)

func collatzSize(n int) int {
    if n == 1 {
        return 1
    }

    if n % 2 == 0 {
        n = n / 2
    } else {
        n = 3*n+1
    }

    return 1+collatzSize(n)
}

func main() {
    s := 0
    n := 0

    for i := 2; i < 1e6; i++ {
        ts := collatzSize(i)
        if ts > s {
            s = ts
            n = i
        }
    }

    fmt.Printf("n: %d, size: %d\n", n, s)
}
