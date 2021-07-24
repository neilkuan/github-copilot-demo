package main

import "fmt"

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]). Return the running sum of nums.

func runningSum(nums []int) []int {
	var sum int
	var result []int
	for _, num := range nums {
		sum += num
		result = append(result, sum)
	}
	return result
}

func main() {
	k := runningSum([]int{1, 2, 3, 4})
	fmt.Println(k)
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	var result []int
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-num {
				result = append(result, i, j)
			}
		}
	}
	return result
}
