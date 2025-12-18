package easy

import "fmt"

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, val := range nums {
		if indx, exists := mp[nums[i]-target]; exists {
			return []int{indx, i}
		}
		mp[val] = i
	}
	return nil
}

func TwoSum() {
	nums := []int{2, 6, 11, 3, 15, 4, 7}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)

}
