// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

package main

func test026() {

}

func removeDuplicates(nums []int) int {
	a := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			a++
			nums[a] = nums[i]
		}
	}
	a = a + 1
	return a
}
