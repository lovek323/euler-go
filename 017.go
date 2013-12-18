package main

import (
    "fmt"
    "strings"
)

func numberToWords(x int) string {
    switch {
    case x == 0: return "zero"
    case x == 1: return "one"
    case x == 2: return "two"
    case x == 3: return "three"
    case x == 4: return "four"
    case x == 5: return "five"
    case x == 6: return "six"
    case x == 7: return "seven"
    case x == 8: return "eight"
    case x == 9: return "nine"
    case x == 10: return "ten"
    case x == 11: return "eleven"
    case x == 12: return "twelve"
    case x == 13: return "thirteen"
    case x == 14: return "fourteen"
    case x == 15: return "fifteen"
    case x == 16: return "sixteen"
    case x == 17: return "seventeen"
    case x == 18: return "eighteen"
    case x == 19: return "nineteen"
    case x == 20: return "twenty"
    case x == 30: return "thirty"
    case x == 40: return "forty"
    case x == 50: return "fifty"
    case x == 60: return "sixty"
    case x == 70: return "seventy"
    case x == 80: return "eighty"
    case x == 90: return "ninety"
    case x % 1000 == 0: return numberToWords(x / 1000) + " thousand"
    case x % 100 == 0: return numberToWords(x / 100) + " hundred"
    case x < 100: return numberToWords(x - x % 10) + "-" + numberToWords(x % 10)
    case x < 1000: return numberToWords(x - x % 100) + " and " + numberToWords(x % 100)
    }

    return ""
}

func main() {
    s := ""

    for i := 1; i <= 1000; i++ {
        s += numberToWords(i)
    }

    s = strings.Replace(s, " ", "", -1)
    s = strings.Replace(s, "-", "", -1)

    fmt.Printf("Result: %d\n", len(s))
}
