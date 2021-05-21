package solve

func PotentialZeroPositions(zeroPosition Position, puzzleSize int) []Position {
	var positions []Position

	if zeroPosition.row > 0 {
		positions = append(
			positions,
			Position{zeroPosition.row - 1, zeroPosition.column},
		)
	}

	if zeroPosition.row < puzzleSize-1 {
		positions = append(
			positions,
			Position{zeroPosition.row + 1, zeroPosition.column},
		)
	}

	if zeroPosition.column > 0 {
		positions = append(
			positions,
			Position{zeroPosition.row, zeroPosition.column - 1},
		)
	}

	if zeroPosition.column < puzzleSize-1 {
		positions = append(
			positions,
			Position{zeroPosition.row, zeroPosition.column + 1},
		)
	}

	return positions
}

func ZeroPosition(puzzle [][]int) Position {
	for i := range puzzle {
		for j, pieceNumber := range puzzle[i] {
			if pieceNumber == 0 {
				return Position{i, j}
			}
		}
	}

	return Position{}
}

func CopyPuzzle(puzzle [][]int) [][]int {
	puzzleCopy := make([][]int, len(puzzle))
	for i := range puzzle {
		puzzleCopy[i] = make([]int, len(puzzle[i]))
		copy(puzzleCopy[i], puzzle[i])
	}

	return puzzleCopy
}

func SwapPuzzlePieces(puzzle [][]int, p1 Position, p2 Position) [][]int {
	newPuzzle := CopyPuzzle(puzzle)

	tmp := newPuzzle[p1.row][p1.column]
	newPuzzle[p1.row][p1.column] = newPuzzle[p2.row][p2.column]
	newPuzzle[p2.row][p2.column] = tmp

	return newPuzzle
}

func Neighbors(node *Node, solvedPiecePositions map[int]Position) []*Node {
	puzzleSize := len(node.puzzle)
	potentialZeroPositions := PotentialZeroPositions(node.zeroPosition, puzzleSize)
	neighbors := make([]*Node, len(potentialZeroPositions))

	for i, position := range potentialZeroPositions {
		newPuzzle := SwapPuzzlePieces(node.puzzle, node.zeroPosition, position)
		cost := node.cost + 1

		neighbors[i] = &Node{
			newPuzzle,
			cost, 0, 0,
			position,
			node,
		}
	}

	return neighbors
}
