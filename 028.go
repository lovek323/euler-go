package main

import (
    "fmt"
    "math"
)

const SIZE int = 1001

// there is a much nicer way of doing this, by understanding that there are
// n/2+1 rings
//
// with the 5x5 spiral:
//
// r0 = 1
//
// r1 = 7 8 9
//      6   2
//      5 4 3
//
// r2 = 21 22 23 24 25
//      20          10
//      19          11
//      18          12
//      17 16 15 14 13
//
// but the brute-force solution was good enough
func calculateSum() int {
    var r [SIZE][SIZE]int
    // start at the rightmost corner and move inwards
    x, y, s := SIZE-1, 0, 0
    d := "left"

    for i := 0; i < len(r); i++ {
        for j := 0; j < len(r[i]); j++ {
            r[i][j] = 0;
        }
    }

    N: for n := SIZE*SIZE; n >= 1; n-- {
        r[y][x] = n

        // are we on a diagonal?
        if x == y || int(math.Abs(float64(x-(SIZE-1)))) == y {
            s += n
        }

        // find next coordinates

        // can we continue to go left
        if d == "left" && x > 0 && r[y][x-1] == 0 {
            x--
            continue N
        }

        // can we continue to go down
        if d == "down" && y < SIZE-1 && r[y+1][x] == 0 {
            y++
            continue N
        }

        // can we continue to go right
        if d == "right" && x < SIZE-1 && r[y][x+1] == 0 {
            x++
            continue N
        }

        // can we continue to go up
        if d == "up" && y > 0 && r[y-1][x] == 0 {
            y--
            continue N
        }

        // can we go left
        if x > 0 && r[y][x-1] == 0 {
            d = "left"
            x--
            continue N
        }

        // can we go down
        if y < SIZE-1 && r[y+1][x] == 0 {
            d = "down"
            y++
            continue N
        }

        // can we go right
        if x < SIZE-1 && r[y][x+1] == 0 {
            d = "right"
            x++
            continue N
        }

        // can we go up
        if y > 0 && r[y-1][x] == 0 {
            d = "up"
            y--
            continue N
        }
    }

    // fmt.Printf("%v\n", r)

    return s
}

func main() {
    s := calculateSum()
    fmt.Printf("%d\n", s)
}
