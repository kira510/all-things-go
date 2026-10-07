package main

func moveZeroes(nums []int) {
	currentZeroPos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != currentZeroPos {
				nums[currentZeroPos] = nums[i]
				nums[i] = 0
			}

			currentZeroPos++
		}
	}
}

/*
https://leetcode.com/problems/move-zeroes/submissions/2164314062/
*/
