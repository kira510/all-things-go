package main

func singleNumber(nums []int) int {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for key, value := range countMap {
		if value == 1 {
			return key
		}
	}

	return -1
}

//https://leetcode.com/problems/single-number/
