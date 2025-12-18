package main

import "fmt"

func hasTwoSum(nums []int, k int) map[int]int {
	compres := make(map[int]int)
	var comp int
	for _, num := range nums {
		comp = k - num
		if _, ok := compres[comp]; ok {
			return map[int]int{num: comp}
		}
		compres[num] = comp

	}
	return map[int]int{}
}
func main() {
	nums := []int{10, 15, 3, 7, 8, 2, 9, 14}
	k := 17

	res := hasTwoSum(nums, k)
	fmt.Println(res)
}
