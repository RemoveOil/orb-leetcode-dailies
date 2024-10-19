package countmaxbitwisesubsets

const QUESTION_NO = 2044

func CountMaxOrSubsets(nums []int) int {
	maxOr := orEveryThing(nums)

	res := 0
	traverseSubsets(nums, 0, 0, maxOr, &res)
	return res
}

func traverseSubsets(nums []int, idx, cumulativeOr, maxOr int, res *int) {
	if idx == len(nums) {
		if cumulativeOr == maxOr {
			*res += 1
		}
		return
	}
	traverseSubsets(nums, idx+1, cumulativeOr, maxOr, res)           // Do not include nums[idx]
	traverseSubsets(nums, idx+1, cumulativeOr|nums[idx], maxOr, res) // Include nums[idx]
}

func orEveryThing(nums []int) int {
	var res int
	for _, num := range nums {
		res |= num
	}
	return res
}
