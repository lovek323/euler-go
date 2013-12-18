package main

import "fmt"

func main() {
    sum := 0
    prev := [...]int{1, 1}

    for cur := 1; cur <= 4e6; cur = prev[0] + prev[1] {
        if cur % 2 == 0 {
            sum += cur
        }
        prev[0] = prev[1]
        prev[1] = cur
    }

    fmt.Printf("Sum: %d\n", sum)
}
