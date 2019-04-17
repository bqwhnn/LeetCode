// 给定一个正整数，输出它的补数。补数是对该数的二进制表示取反。

// 注意:

// 给定的整数保证在32位带符号整数的范围内。
// 你可以假定二进制数不包含前导零位。
// 示例 1:

// 输入: 5
// 输出: 2
// 解释: 5的二进制表示为101（没有前导零位），其补数为010。所以你需要输出2。
// 示例 2:

// 输入: 1
// 输出: 0
// 解释: 1的二进制表示为1（没有前导零位），其补数为0。所以你需要输出0。

package main

func test476() {

}

func findComplement(num int) int {
    res := 0
    arr := []int{}
    
    for (num != 0) {
        arr = append(arr, num&1)
        num >>= 1
    }
    
    for i := len(arr)-1; i >=0; i-- {
        res <<= 1
        res += arr[i]^1
    }
    return res
}

// 取得 num 位数都为 1 的二进制 n，在异或即可按位取反
func findComplement2(num int) int {
    temp := num
    n := 0
    for temp > 0 {
        temp >>= 1
        n = (n<<1) + 1
    }
    return num ^ n
}