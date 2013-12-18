package main

import (
    "fmt"
    "strconv"
)

func isPalendromic(x int) bool {
    xs := strconv.Itoa(x)
    sx := make([]rune, len(xs))

    for i := 1; i <= len(xs); i++ {
        sx[i-1] = rune(xs[len(xs)-i]);
    }

    return string(sx) == xs;
}

func main() {
    var result int;

    for i := 100; i <= 999; i++ {
        for j := 100; j <= 999; j++ {
            if isPalendromic(i * j) && i * j > result {
                result = i * j
            }
        }
    }

    fmt.Printf("Result: %d\n", result)
}
