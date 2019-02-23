package main

import "sort"

func main() {

}

func sortedSquares(A []int) []int {
    r := []int{}
    for _, v := range A {
        r = append(r, v*v)
    }
    sort.Ints(r)
    return r
}