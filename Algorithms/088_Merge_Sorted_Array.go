// 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
// 
// 说明:
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

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