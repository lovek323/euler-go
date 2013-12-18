package main

import (
    "fmt"
    "math/big"
    "strconv"
)

func factorial(n *big.Int) *big.Int {
    if n.Cmp(big.NewInt(0)) == -1 {
        return big.NewInt(1)
    }

    if n.Cmp(big.NewInt(0)) == 0 {
        return big.NewInt(1)
    }

    r := new(big.Int)
    r.Set(n)
    r.Mul(r, factorial(n.Sub(n, big.NewInt(1))))

    return r
}

func main() {
    nf := factorial(big.NewInt(100))
    nfs := nf.String()
    s := 0

    for i := 0; i < len(nfs); i++ {
        t, _ := strconv.Atoi(nfs[i:i+1])
        s += t
    }

    fmt.Printf("Result: %d\n", s)
}
