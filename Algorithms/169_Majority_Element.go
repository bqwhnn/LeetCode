// 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 
// 你可以假设数组是非空的，并且给定的数组总是存在众数。

package main

import (
    // "fmt"
)

func test169() {
	
}

func majorityElement(nums []int) int {
    ret := nums[0]
    cnt := 0
    for _, num := range nums {
        if cnt == 0 {
            ret = num
        }
        if num == ret {
            cnt++
        } else {
            cnt--
        }
    }
    return ret
}