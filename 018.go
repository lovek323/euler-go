package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

const TS string =
"75\n"+
"95 64\n"+
"17 47 82\n"+
"18 35 87 10\n"+
"20 04 82 47 65\n"+
"19 01 23 75 03 34\n"+
"88 02 77 73 07 63 67\n"+
"99 65 04 28 06 16 70 92\n"+
"41 41 26 56 83 40 80 70 33\n"+
"41 48 72 33 47 32 37 16 94 29\n"+
"53 71 44 65 25 43 91 52 97 51 14\n"+
"70 11 33 28 77 73 17 78 39 68 17 57\n"+
"91 71 52 38 17 14 91 43 58 50 27 29 48\n"+
"63 66 04 68 89 53 67 30 73 16 69 87 40 31\n"+
"04 62 98 27 23 09 70 98 73 93 38 53 60 04 23"

func getTriangleMax(t [][]int) int {
    if len(t) == 1 {
        return t[0][0]
    }

    // start at the second-last row and determine the maximum values for each
    // entry

    // the first array is the array of rows, so the row count = len(t) and the
    // column count of a particular row is len(t[row])

    // row0 will be replaced with the maximum values
    row0 := t[len(t)-2]
    row1 := t[len(t)-1]

    for i := 0; i < len(row0); i++ {
        // we use -1 to indicate that a cell is empty
        if row0[i] == -1 {
            break
        }

        // for each element, determine the maximum value
        // element row0[i] can go to row1[i] and row1[i+1]
        row0[i] = int(math.Max(float64(row0[i]+row1[i]),float64(row0[i]+row1[i+1])))
    }

    // remove the last row and re-calculate
    t_ := t[0:len(t)-1]

    return getTriangleMax(t_)
}

type Triangle [][]int

func main() {
    ts := strings.Split(TS, "\n")
    rn := len(ts)
    cn := len(strings.Split(ts[rn-1], " "))

    var t Triangle

    for i := 0; i < rn; i++ {
        rows := strings.Split(ts[i], " ")
        var rowi []int
        for j := 0; j < len(rows); j++ {
            v, _ := strconv.Atoi(rows[j])
            rowi = append(rowi, v)
        }
        for j := len(rows); j < cn; j++ {
            rowi = append(rowi, -1)
        }
        t = append(t, rowi)
    }

    fmt.Printf("Result: %d\n", getTriangleMax(t))
}
