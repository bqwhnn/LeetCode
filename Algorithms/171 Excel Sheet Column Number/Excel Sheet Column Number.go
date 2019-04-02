package main

func main() {

}

// 26 进制
func titleToNumber(s string) int {
    n := 0
    for _, char := range s {
        n = n * 26 + int(char) - 'A'
    }
    return n
}