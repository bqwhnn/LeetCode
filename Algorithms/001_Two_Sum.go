package main

import (
    "fmt"
)

func test001() {
    nums := []int{2, 3, 6, 11}
    fmt.Println(twoSum(nums, 9))
    fmt.Println(twoSum1(nums, 8))
}

// 暴力法
func twoSum(nums []int, target int) []int {
    for i, j := range nums{
        for k, l := range nums {
            if i != k && j + l == target {
                return []int{i, k}
            }
        }
    }
    return []int{}
}

// 利用 go 自带的 hash map
func twoSum1(nums []int, target int) []int {
    m := make(map[int]int)
    for k, v := range nums {
        t := target - v
        if i, ok := m[t]; ok {
            return []int{i, k}
        }
        m[v] = k
    }
    return []int{}
}