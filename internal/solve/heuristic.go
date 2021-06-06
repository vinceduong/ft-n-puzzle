package solve

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(row int, column int, pos Position) int {
	return Abs(row-pos.row) + Abs(column-pos.column)
}

func Heuristic(puzzle [][]int, solved map[int]Position) int {
	heuristic := 0

	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			distance := manhattanDistance(i, j, solved[pieceNumber])
			heuristic += distance
		}
	}

	return heuristic
}
