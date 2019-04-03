package main

import (
    "fmt"
)

func test136() {
    nums := []int{1, 2, 2, 1, 5, 6, 5}
    fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
    a := 0
    for _, num := range nums {
        a ^= num
    }
    return a
}