package lib

import "math"

func IsPrime64(x int64) bool {
    var i int64

    for i = 2; i <= int64(math.Sqrt(float64(x))); i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}

func IsPrime(x int) bool {
    for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}
