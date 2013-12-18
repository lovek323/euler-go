package main

import (
    "fmt"
    "math/big"
    "strconv"
)

func main() {
    n := new(big.Int).Exp(big.NewInt(2),big.NewInt(1000),nil)
    ns := n.String()
    s := 0

    for i := 0; i < len(ns); i++ {
        t, _ := strconv.Atoi(ns[i:i+1])
        s += t
    }

    fmt.Printf("Result: %d\n", s)
}
