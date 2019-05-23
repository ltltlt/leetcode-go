package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	type Direction int
	const (
		right Direction = iota
		down
		left
		up
	)
	var (
		rows       = len(matrix)
		cols       = len(matrix[0])
		curX, curY int
		results    []int
		direction  = right
	)
	var scanned = make([][]bool, rows)
	for i := range scanned {
		scanned[i] = make([]bool, cols)
	}
	for !scanned[curX][curY] {
		results = append(results, matrix[curX][curY])
		scanned[curX][curY] = true
		for i := 0; i < 4; i++ {
			var tmpX, tmpY int
			switch direction {
			case right:
				tmpX, tmpY = curX, curY+1
			case down:
				tmpX, tmpY = curX+1, curY
			case left:
				tmpX, tmpY = curX, curY-1
			case up:
				tmpX, tmpY = curX-1, curY
			}
			if tmpX >= 0 && tmpX < rows && tmpY >= 0 && tmpY < cols && !scanned[tmpX][tmpY] {
				curX, curY = tmpX, tmpY
				break
			}
			direction = (direction + 1) % 4
		}
	}
	return results
}
