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