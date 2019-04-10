// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 注意：给定 n 是一个正整数。

package main

func test070() {

}

// 斐波那契数列
// 首先考虑第一步的走法：第一步可以走1级台阶，也可以走2级台阶。
// 那么，设S(n)表示走n级台阶的走法数量，
// 如果第一步走1级台阶，剩下的台阶数为n-1，也就是说这种情况下的走法是相当于S(n-1)；
// 同理，如果第一步走2级台阶，剩下的台阶数为n-2，这样的走法相当于S(n-2)；于是，得出递推公式：
// S(n) = S(n-1) + S(n-2);
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    a := 1
    b := 2
    for i := 3; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}