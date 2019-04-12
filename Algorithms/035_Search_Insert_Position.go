// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 你可以假设数组中无重复元素。

package main

func test035() {

}

// o(n)
func searchInsert(nums []int, target int) int {
    if len(nums) <= 0 || target < nums[0] {
        return 0
    }
    for k, num := range nums {
        if num >= target {
            return k
        } 
    }
    return len(nums)
}

// o(logn)
func searchInsert2(nums []int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] > target {
            r = mid
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            return mid
        }
    }
    return l
}