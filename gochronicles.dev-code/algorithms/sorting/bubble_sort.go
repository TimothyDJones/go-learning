package sorting

func BubbleSort(nums []int) []int {
	n := len(nums)

	for i := 1; i <= (n - 1); i++ {
		for j := (n - 1); j >= i; j-- {
			if nums[j] < nums[j - 1] {
				nums[j], nums[j - 1] = nums[j - 1], nums[j]
			}
		}
	}

	return nums
}
