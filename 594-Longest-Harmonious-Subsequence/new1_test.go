package leetcode594

func findLHS2(nums []int) int {
	var counter = make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num] += 1
	}

	var max int
	for num, c := range counter {
		if c1, ok := counter[num+1]; ok {
			count := c + c1
			if max < count {
				max = count
			}
		}
	}
	return max
}
