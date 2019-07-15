package leetcode

import "math"

// EMPTY 空的房间
const EMPTY = math.MaxInt32

// GATE 门
const GATE = 0

// DIRECTIONS 不知道
var DIRECTIONS = [][]int{
	[]int{1, 0},
	[]int{-1, 0},
	[]int{0, 1},
	[]int{0, -1},
}

func wallsAndGates(rooms [][]int) {
	m := len(rooms)

	if m == 0 {
		return
	}

	n := len(rooms[0])

	var queue [][]int

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if rooms[row][col] == GATE {
				queue = append(queue, []int{row, col})
			}
		}
	}

	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		row, col := point[0], point[1]
		for _, direction := range DIRECTIONS {
			r := row + direction[0]
			c := col + direction[1]
			if r < 0 || c < 0 || r >= m || c >= n || rooms[r][c] != EMPTY {
				continue
			}
			rooms[r][c] = rooms[row][col] + 1
			queue = append(queue, []int{r, c})
		}
	}
}
