package main

import (
    "fmt"
    "./lib"
)

var triangles []int

func nextTriangle() int {
    triangles = append(triangles, triangles[len(triangles)-1]+len(triangles)+1)

    return triangles[len(triangles)-1]
}

func main() {
    triangles = append(triangles, 1)
    triangle := 1
    divisors := 0

    for ; divisors < 500; {
        triangle = nextTriangle()
        divisors = lib.NumDivisors(triangle)
    }

    fmt.Printf("Triangle number: %d, %d divisors\n", triangle, divisors)
}
