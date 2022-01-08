package QuickSort

import "math/rand"

func QuickSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	begin, end := 0, length-1
	pivot := rand.Intn(length) // 选择用于划分的元素
	nums[pivot], nums[0] = nums[0], nums[pivot] // 将元素交换到第一个位置
	for begin < end {
		for begin < end && nums[end] >= nums[0] {
			end--
		}
		for begin < end && nums[begin] <= nums[0] {
			begin++
		}
		nums[begin], nums[end] = nums[end], nums[begin]
	}
	nums[begin], nums[0] = nums[0], nums[begin]
	QuickSort(nums[:begin])
	QuickSort(nums[begin+1:])
	return nums
}
