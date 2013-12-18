package main

import (
    "fmt"
    "math/big"
    "strconv"
    "./lib"
)

func main() {
    nf := lib.FactorialBig(big.NewInt(100))
    nfs := nf.String()
    s := 0

    for i := 0; i < len(nfs); i++ {
        t, _ := strconv.Atoi(nfs[i:i+1])
        s += t
    }

    fmt.Printf("Result: %d\n", s)
}
