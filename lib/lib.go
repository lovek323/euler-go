package lib

import (
    "math"
    "math/big"
)

func IsPrime64(x int64) bool {
    var i int64

    for i = 2; i <= int64(math.Floor(math.Sqrt(float64(x)))); i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}

func IsPrime(x int) bool {
    for i := 2; i <= int(math.Floor(math.Sqrt(float64(x)))); i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}

func NumDivisors(x int) int {
    c := 2 // 1 and itself

    for i := 2; i <= int(math.Floor(math.Sqrt(float64(x)))); i++ {
        if x % i == 0 {
            c += 2
        }
    }

    return c
}

func GetProperDivisors(x int) []int {
    d := []int{1}

    for i := 2; i < int(math.Floor(math.Sqrt(float64(x))))+1; i++ {
        if x % i == 0 {
            d = append(d, i)

            if (x / i != i) {
                d = append(d, x / i)
            }
        }
    }

    return d
}

func FactorialBig(n *big.Int) *big.Int {
    if n.Cmp(big.NewInt(0)) == -1 {
        return big.NewInt(1)
    }

    if n.Cmp(big.NewInt(0)) == 0 {
        return big.NewInt(1)
    }

    r := new(big.Int)
    r.Set(n)
    r.Mul(r, FactorialBig(n.Sub(n, big.NewInt(1))))

    return r
}

func Factorial(n int) int {
    if n < 0 || n == 0 {
        return 1
    }

    return n * Factorial(n-1)
}
