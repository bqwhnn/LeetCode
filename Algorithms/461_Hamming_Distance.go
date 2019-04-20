// 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。

// 给出两个整数 x 和 y，计算它们之间的汉明距离。

// 注意：
// 0 ≤ x, y < 231.

package main

func test461() {

}

func hammingDistance(x int, y int) int {
    n := 0
    for (x != 0 || y != 0) {
        if (x & 1 != y & 1) {
            n++
        }
        x >>= 1
        y >>= 1
    }
    return n
}