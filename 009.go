package main

import (
    "fmt"
)

func nextTriplet() (int, int, int) {
    for a := 1; a <= 1000; a++ {
        for b := a+1; b <= (1000-a); b++ {
            for c := b+1; c <= (1000-a-b); c++ {
                if a + b + c == 1000 {
                    if a*a+b*b == c*c {
                        return a, b, c
                    }
                }
            }
        }
    }

    return 0, 0, 0
}

func main() {
    a, b, c := nextTriplet()
    fmt.Printf("Result: %d\n", a*b*c)
}
