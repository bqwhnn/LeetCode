// 给定一个非负整数数组 A，返回一个由 A 的所有偶数元素组成的数组，后面跟 A 的所有奇数元素。

// 你可以返回满足此条件的任何数组作为答案。

package main

func test905() {
	
}

func sortArrayByParity(A []int) []int {
    a := 0;
    b := len(A) - 1
    for a < b {
        if A[a] & 1 == 0 {
            a++
            continue
        }
        if A[b] & 1 == 1 {
            b--
            continue
        }
        A[a], A[b] = A[b], A[a]
        a++
        b--
    }
    return A
}