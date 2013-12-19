package main

import (
    "fmt"
    "math/big"
    "strings"
)

func getRecurringLengthRat(x *big.Rat) int {
    xs := strings.Trim(x.FloatString(3000), "0")

    return getRecurringLength(xs)
}

func getRecurringLength(xs string) int {
    lr := 0

    // cut off leading decimal point
    xs = xs[1:]

    if len(xs) < 3000 {
        return 0
    }

    for i := 1; i < len(xs)/3; i++ {
        // test to see if a string of length i is repeated
        J: for j := 0; j < len(xs)-3*i-1; j++ {
            // if string of length i at position j is repeated at least twice,
            // we have the largest recursion
            start := j
            end := start+i
            xsi := xs[start:end]

            start = end
            end = start+i

            for ; end < len(xs); {
                xsii := xs[start:end]
                if xsi != xsii {
                    continue J
                }
                start = end
                end = start+i
            }

            lr = i
            break
        }

        if lr != 0 {
            break
        }
    }

    return lr
}

func main() {
    var d, ld int64
    ll := 0

    for d = 1; d < 1000; d++ {
        r := big.NewRat(1,d)
        l := getRecurringLengthRat(r)
        if l > ll {
            ld = d
            ll = l
        }
    }

    fmt.Printf("Result: %d\n", ld)
}
