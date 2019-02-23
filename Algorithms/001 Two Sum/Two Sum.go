package main

func main() {
	
}

func twoSum(nums []int, target int) []int {
    for i, j := range nums{
        for k, l := range nums {
            if i != k && j + l == target {
                return []int{i, k}
            }
        }
    }
    return []int{}
}

func twoSum_1(nums []int, target int) []int {
    m := make(map[int]int)
    for k, v := range nums {
        t := target - v
        if i, ok := m[t]; ok {
            return []int{i, k}
        }
        m[v] = k
    }
    return []int{}
}