// 给定一个整数 n，返回 n! 结果尾数中零的数量。

package main

func test172() {

}

func trailingZeroes(n int) int {
	i := 0
	for n >= 5 {
		i += n / 5
		n = n / 5
	}
	return i
}
