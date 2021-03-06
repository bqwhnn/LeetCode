// 实现 strStr() 函数。
//
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
// 当 needle 是空字符串时我们应当返回 0 。

package main

import "fmt"

func test028() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    for i := 0; i < len(haystack); i++ {
        if haystack[i] == needle[0] && i+len(needle) <= len(haystack) {
            tmp := []byte(haystack)[i:i+len(needle)]
            if string(tmp) == needle {
                return i
            }
        }
    }
    return -1
}