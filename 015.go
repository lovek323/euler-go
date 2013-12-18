package main

import (
    "fmt"
    "math/big"
)

func main() {
    // the number of paths to get to the bottom right of an m x n grid is
    // binomial(m+n,n) (need m moves down and n moves right in any order)
    fmt.Printf("Result: %d\n", new(big.Int).Binomial(40,20))
}
