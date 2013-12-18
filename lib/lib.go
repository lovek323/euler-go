package lib

import "math"

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
