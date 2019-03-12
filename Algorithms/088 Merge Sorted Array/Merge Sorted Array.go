package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    total := m+n-1
    m--
    n--
    for i := total; i >= 0; i-- {
        if m < 0 || n < 0 {
            break
        }
        if nums1[m] > nums2[n] {
            nums1[i] = nums1[m]
            m--
        } else {
            nums1[i] = nums2[n]
            n--
        }
    }
    for i := n; i >= 0; i-- {
        nums1[i] = nums2[i]
    }
}