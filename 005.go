package main

import (
    "fmt"
)

func isDivisibleRange(k0 int, k1 int, x int) bool {
    for i := k0; i <= k1; i++ {
        if (x % i != 0) {
            return false;
        }
    }

    return true;
}

func main() {
    for i := 20; ; i += 20 {
        if isDivisibleRange(1, 20, i) {
            fmt.Printf("Result: %d\n", i)
            break
        }
    }
}
