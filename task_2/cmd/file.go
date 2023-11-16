package main

import (
	"fmt"
	"os"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		fmt.Println("Ошибка, ваш массив пустой!")
		os.Exit(1)
	}
	nums = quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(nums []int, start int, last int) []int {
	if start < last {
		var root int
		nums, root = partition(nums, start, last)
		nums = quickSort(nums, start, root-1)
		nums = quickSort(nums, root+1, last)
	}
	return nums
}

func partition(nums []int, start int, last int) ([]int, int) {
	pivot := nums[last]
	i := start
	for j := start; j < last; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[last] = nums[last], nums[i]
	return nums, i
	
}
