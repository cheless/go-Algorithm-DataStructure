package main

import (
	"fmt"
)

func main() {
	t := []int{4,7,5,9,2}
	t = mergeSort(t)
	fmt.Println(t)
}

func mergeSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	left, right := mergeSort(nums[0:l/2]), mergeSort(nums[l/2:])
	rst := make([]int, 0, l)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			rst = append(rst, left[i])
			i++
		} else {
			rst = append(rst, right[j])
			j++
		}
	}
	if i == len(left) {
		rst = append(rst, right[j:]...)
	} else {
		rst = append(rst, left[i:]...)
	}
	return rst
}
