package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func read() []string {
    b, _ := ioutil.ReadFile("data/022.txt")
    s := string(b)

    return strings.Split(strings.Replace(s, "\"", "", -1), ",")
}

// simple implementation of quick sort (because it's fun)
func sort(s []string) []string {
    if len(s) == 0 {
        return nil
    }

    first := s[0]
    s = s[1:]

    // why use my own function instead of s < first and s > first?
    // again, because it's fun (and my implementation only adds about 0.1 s)
    lf := func (s string) bool { return compareString(s, first) == -1 }
    gf := func (s string) bool { return compareString(s, first) == 1 }

    l := sort(filter(s, lf))
    g := sort(filter(s, gf))

    var r []string
    for _, si := range l {
        r = append(r, si)
    }
    r = append(r, first)
    for _, si := range g {
        r = append(r, si)
    }

    return r
}

func filter(s []string, f (func(s string) bool)) []string {
    var r []string
    for _, p := range s {
        if f(p) {
            r = append(r, p)
        }
    }
    return r
}

func compareString(x string, y string) int {
    // the first letter is the same, remove it and move on to the next
    c := compareByte(x[0], y[0]);

    if c == 0 {
        if len(x) == 1 {
            if len(y) == 1 {
                return 0
            }

            // x is shorter than y and thus should come first
            return -1
        }

        if len(y) == 1 {
            // x is longer than y and thus should come second
            return 1
        }

        return compareString(x[1:len(x)], y[1:len(y)])
    }

    return c
}

func compareByte(x uint8, y uint8) int {
    if x == y {
        return 0
    }
    if x > y {
        return 1
    }
    return -1
}

func getScore(x string) int {
    s := 0
    for i := 0; i < len(x); i++ {
        s += getScore8(x[i])
    }
    return s
}

func getScore8(x uint8) int {
    return int(x - uint8('A') + 1)
}

func main() {
    sr := read()
    ss := sort(sr)
    s := 0

    for i := 0; i < len(ss); i++ {
        s += getScore(ss[i])*(i+1)
    }

    fmt.Printf("Result: %d\n", s)
}
