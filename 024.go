package main

import (
    "fmt"
    "strconv"
    "./lib"
)

// there are k! permutations of numbers 0-(k-1), i.e., we have k = 10 as we want
// to order the numbers 0-9

// 0 is the first in the lexical order
// the last 9 digits can be ordered in 9! = 362880 ways, thus the first 362880
// permutations start with 0

// the permutations 9!+1 to 2*9! start with a 1
// the permutations 2*9!+1 to 3*9! start with a 2
// thus, the millionth interval must start with a 2

/**
 * n - the number of the permutation we wish to find
 * k - the number of items in our list of permutations
 *
 * returns the segment in which the nth permutation sits and the first
 * permutation in that segment
 */
func findNthPermutationSegment(n, k int) (int, int) {
    fact := lib.Factorial(k-1)
    var seg int

    for seg = 0; seg < k; seg++ {
        if seg*fact <= n && (seg+1)*fact >= n {
            break;
        }
    }

    return seg, seg*fact
}

func removeNthElement(n int, x []int) []int {
   k := make([]int, len(x)-1)

   for i := 0; i < len(x); i++ {
       if (i == n) {
           continue
       }
       if (i > n) {
           k[i-1] = x[i]
       } else {
           k[i] = x[i]
       }
   }

   return k
}

func findNthPermutation(n int, k int) int {
    var nums []int
    for i := 0; i < k; i++ { nums = append(nums, i) }

    nthPermutation := ""

    for ; len(nums) > 1; {
        seg, segStart := findNthPermutationSegment(n, len(nums))
        nthPermutation += strconv.Itoa(nums[seg])
        n -= segStart
        nums = removeNthElement(seg, nums)
    }

    nthPermutation += strconv.Itoa(nums[0])
    r, _ := strconv.Atoi(nthPermutation)

    return r
}

func main () {
    fmt.Printf("Result: %d\n", findNthPermutation(1e6, 10))
}
