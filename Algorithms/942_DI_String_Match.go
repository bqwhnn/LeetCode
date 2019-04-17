// 给定只含 "I"（增大）或 "D"（减小）的字符串 S ，令 N = S.length。

// 返回 [0, 1, ..., N] 的任意排列 A 使得对于所有 i = 0, ..., N-1，都有：

// 如果 S[i] == "I"，那么 A[i] < A[i+1]
// 如果 S[i] == "D"，那么 A[i] > A[i+1]

package main

func test942() {
	
}

func diStringMatch(S string) []int {
    min := 0
    max := len(S)
    res := make([]int, 0)
    
    for _, c := range S {
        if c == 'I' {
            res = append(res, min)
            min++
        } else {
            res = append(res, max)
            max--
        }
    } 
    res = append(res, min)
    return res
}

// 小优化，少去了 append 的消耗
func diStringMatch2(S string) []int {
    min := 0
    max := len(S)
    res := make([]int, len(S)+1)
    
    for i, c := range S {
        if c == 'I' {
            res[i] = min
            min++
        } else {
            res[i] = max
            max--
        }
    }
    res[len(S)] = min
    return res
}