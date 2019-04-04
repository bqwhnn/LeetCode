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