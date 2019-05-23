package nqueens

import (
	"fmt"
	"strings"
	"testing"
)

// board: a list
// example: 8x8 board is a list which length is 8, each element is a number indicate
// which column the queen puts

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// each call should return a board, not a partial board
func recursive(n int, positions []int) [][]int {
	if len(positions) >= n {
		return nil
	}
	nextp := len(positions)
	var results [][]int
	for pos := 0; pos < n; pos++ {
		ok := true
		for i, pos1 := range positions {
			off := pos - pos1
			if off == 0 || abs(off) == abs(i-nextp) {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		npositions := make([]int, len(positions)+1)
		copy(npositions, positions)
		npositions[len(npositions)-1] = pos
		// last
		if len(positions) == n-1 {
			results = append(results, npositions)
		} else {
			results = append(results, recursive(n, npositions)...)
		}
	}
	return results
}

// [0, 1] -> ["Q.", ".Q"]
func dump(solutions [][]int) [][]string {
	var results [][]string
	for _, solution := range solutions {
		var result []string
		for _, pos := range solution {
			result = append(result,
				strings.Repeat(".", pos)+"Q"+
					strings.Repeat(".", len(solution)-pos-1))
		}
		results = append(results, result)
	}
	return results
}

func solveNQueens(n int) [][]string {
	solutions := recursive(n, nil)
	return dump(solutions)
}

func TestSolveNQueens(t *testing.T) {
	boards := solveNQueens(4)
	fmt.Println("total solutions:", len(boards))
	for _, board := range boards {
		fmt.Println("------------")
		for _, row := range board {
			fmt.Println(row)
		}
		fmt.Println("------------")
	}
	t.Fail()
}
