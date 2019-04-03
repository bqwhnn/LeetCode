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