package main

import (
    "fmt"
)

func test088() {
    nums1 := []int{1, 2, 3, 0, 0, 0}
    nums2 := []int{2, 5, 6}
    merge(nums1, 3, nums2, 3)
    fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    total := m+n-1
    m--
    n--
    for i := total; i >= 0; i-- {
        if m < 0 || n < 0 {
            break
        }
        if nums1[m] > nums2[n] {
            nums1[i] = nums1[m]
            m--
        } else {
            nums1[i] = nums2[n]
            n--
        }
    }
    for i := n; i >= 0; i-- {
        nums1[i] = nums2[i]
    }
}