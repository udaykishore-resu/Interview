package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Given Array is: ", nums)
	res := getProdExceptSelf(nums)
	fmt.Println("Result array is :", res)
}

func getProdExceptSelf(nums [5]int) []int {
	length := len(nums)
	res := make([]int, length)

	l := 1
	for i, val := range nums {
		res[i] = 1
		res[i] *= l
		l *= val
	}

	r := 1
	for i := length - 1; i >= 0; i-- {
		res[i] *= r
		r *= nums[i]
	}
	return res
}
