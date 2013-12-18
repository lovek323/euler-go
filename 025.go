package main

import (
    "fmt"
    "./lib"
)

func find(i int) int {
    fi := lib.GetNthFibonacci(i)
    fis := fi.String()
    fi1 := lib.GetNthFibonacci(i-1)
    fi1s := fi1.String()

    if len(fis) > 1000 {
        // we need to go back
        return find(i-1)
    } else if len(fis) < 1000 {
        // we need to go forward
        return find(i+10)
    } else if len(fis) == 1000 && len(fi1s) == 1000 {
        // we need to go back
        return find(i-1)
    }

    return i
}

func main() {
    fmt.Printf("Result: %d\n", find(1000))
}
