package array

import "sort"

// leetcode 15. 3Sum
// https://leetcode.com/problems/3sum/
// First approach is to fix one number, and then apply two sum solution
// This won't remove duplicates, so the way to to this is to sort array first.
// On sorted array we can skip numbers we've already checked, which will produce only unique
// triplets.

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums) - 2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for k > j && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
		for i < len(nums) - 2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
