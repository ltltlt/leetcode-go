/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	results := make([][]int, numRows)
	results[0] = []int{1}
	for i := 1; i < numRows; i++ {
		result := make([]int, i+1)
		prev := results[i-1]
		j, elem := 0, 0
		for ; j < len(prev); j++ {
			result[j] = elem + prev[j]
			elem = prev[j]
		}
		result[j] = 1
		results[i] = result
	}

	return results
}

