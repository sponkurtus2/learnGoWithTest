package arraysandslices

func Sum(nums []int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var sum []int

	for _, nums := range numsToSum {
		sum = append(sum, Sum(nums))
	}

	return sum
}

func SumAllTails(numsToSum ...[]int) []int {
	var sum []int

	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sum = append(sum, 0)
		} else {
			tail := nums[len(nums)-1:]
			sum = append(sum, Sum(tail))
		}

	}

	return sum
}
