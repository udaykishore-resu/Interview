package easy

import "fmt"

func PrefixSumArray() {
	nums := []int{1, 8, 6, 7, 4, 5, 3, 2}

	for i, v := range nums {
		if i == 0 {
			continue
		}
		nums[i] = nums[i-1] + v
	}

	fmt.Println(nums)
}
