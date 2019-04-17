// 我们把符合下列属性的数组 A 称作山脉：

// A.length >= 3
// 存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
// 给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。

package main

func test852() {

}

func peakIndexInMountainArray(A []int) int {
    low := 0
    high := len(A) - 1
    mid := 0
    
    for (low < high) {
        mid = low + (high - low) / 2
        if A[mid] > A[mid-1] && A[mid] < A[mid+1] {
            low = mid + 1
        } else if A[mid] < A[mid-1] && A[mid] > A[mid+1] {
            high = mid
        } else {
            return mid
        }
    }
    return -1
}