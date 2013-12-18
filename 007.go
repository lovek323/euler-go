package main

import (
    "fmt"
    "./lib"
)

func main() {
    count := 0
    last := 0

    for i := 2; ; i++ {
        if lib.IsPrime(i) {
            last = i
            count++
        }

        if (count == 10001) {
            break
        }
    }

    fmt.Printf("10,001st prime: %d\n", last)
}
