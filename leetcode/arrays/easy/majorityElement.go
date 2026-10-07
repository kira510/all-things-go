package main

func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}

/*
https://leetcode.com/problems/majority-element/
*/
