package main

import "fmt"

func test171() {
	s := "ACG"
	fmt.Println(titleToNumber(s))
}

// 26 进制
func titleToNumber(s string)int {
    n := 0
    for _, char := range s {
        n = n * 26 + int(char) - 'A' + 1
        fmt.Println(char, n)
    }
    return n
}