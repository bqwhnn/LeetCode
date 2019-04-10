// 给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。

// 水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。

// 反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。

package main

func test832() {

}

func flipAndInvertImage(A [][]int) [][]int {
    for _, arr := range A {
        l := len(arr)
        for i := 0; i < l / 2; i++ {
            arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
        }
        for i := 0; i < l; i++ {
            if arr[i] == 0 {
                arr[i] = 1
            } else {
                arr[i] = 0
            }
        }
    }
    return A
}

func flipAndInvertImage2(A [][]int) [][]int {
	for _, a := range A {
		i, j := 0, len(a) - 1
		for i <= j {
			a[i], a[j] = 1-a[j], 1-a[i]
			i++
			j--
		}
	}
	return A
}