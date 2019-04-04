// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

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