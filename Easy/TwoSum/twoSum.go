package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Print(result)
}

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	result := []int{}

	for i, num := range nums {
		complement := target - num

		if index, exist := myMap[complement]; exist {
			result = append(result, index, i)
			return result
		}
		myMap[num] = i
	}
	return result
}
