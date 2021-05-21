package solve
//import "fmt"
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDistance(row int, column int, pos Position) int {
	return Abs(row-pos.row) + Abs(column-pos.column)
}

func isMissPlaced(row int, column int, pos Position) int {
	if row != pos.row || column != pos.column {
		return 1
	}
	return 0
}

func Heuristic(t string, puzzle [][]int, solvedPiecePositions map[int]Position) int {
	heuristic := 0

	var function func(row int, column int, pos Position) int

	switch t {
	case "manhattan":
		function = ManhattanDistance
	case "mp":
		function = isMissPlaced
	default:
		function = ManhattanDistance
	}

	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			distance := function(i, j, solvedPiecePositions[pieceNumber])
			// distance := isMissPlaced(i, j, solvedPiecePositions[pieceNumber])
			heuristic += distance
		}
	}

//	fmt.Printf("puzzle = %v\n", puzzle)
//	fmt.Printf("heuristic = %v\n", heuristic)
	return heuristic
}
