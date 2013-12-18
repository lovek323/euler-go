package main

import "fmt"

func getDaysSince1900(d int, m int, y int) int {
    n := d
    if m > 1 {
        n += getDaysSince1900(0, m - 1, 0)
        n += getDaysInMonth(m - 1, y)
    }
    if y > 1900 {
        n += getDaysSince1900(0, 0, y - 1)
        n += getDaysInYear(y - 1)
    }

    return n
}

func getDaysInMonth(m int, y int) int {
    switch m {
    case 2:
        if isLeapYear(y) {
            return 29
        }
        return 28
    case 4: case 6: case 9: case 11:
        // 30 days in April, June, September and November
        return 30
    default:
        return 31
    }

    return 0
}

func getDaysInYear(y int) int {
    if isLeapYear(y) {
        return 366;
    }

    return 365;
}

func isLeapYear(y int) bool {
    return y % 4 == 0 && y % 400 != 0
}

func isSunday(d int, m int, y int) bool {
    // 1900-01-01 was a Monday, which we calculate as 1 day after 1900, so
    // any date with the number of days divisible by 7 will be a Sunday
    return getDaysSince1900(d, m, y) % 7 == 0
}

func main() {
    c := 0

    for y := 1901; y <= 2000; y++ {
        for m := 1; m <= 12; m++ {
            if (isSunday(1, m, y)) {
                c++
            }
        }
    }

    fmt.Printf("Result: %d\n", c)
}
