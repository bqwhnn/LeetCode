// 在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。

// 返回重复了 N 次的那个元素。

package main

func test961() {

}

func repeatedNTimes(A []int) int {
    m := make(map[int]bool)
    for _, n := range A {
        if m[n] {
            return n
        }
        m[n] = true
    }
    return -1
}