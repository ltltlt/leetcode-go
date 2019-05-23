package nqueens_test

// Space cost is n + constant, space complexity is O(n)
// we only need one slice which capacity is n, and some temporary variable

func totalNQueens(n int) int {
	return recursive(n, make([]int, 0, n))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// each call should return a board, not a partial board
func recursive(n int, positions []int) int {
	nextp := len(positions)
	if nextp >= n {
		return 1
	}
	var total int
	for pos := int(0); pos < n; pos++ {
		ok := true
		for i, pos1 := range positions {
			off := pos - pos1
			if off == 0 || abs(off) == abs(i-nextp) {
				ok = false
				break
			}
		}
		if ok {
			total += recursive(n, append(positions, pos))
		}
	}
	return total
}
