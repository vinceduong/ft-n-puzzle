package solve

func Solvable(puzzle, solvedPuzzle [][]int) bool {
	inversions, blankPieceRow := Inversions(puzzle, solvedPuzzle)

	/*
			If the puzzle size is even, puzzle instance is solvable if
		  the blank is on an even row counting from the bottom (second-last, fourth-last, etc.)
			and number of inversions is odd.
	*/
	if len(puzzle)%2 == 0 && ((len(puzzle)-1)-blankPieceRow)%2 == 0 {
		return inversions%2 != 0
	}
	/*
		If the puzzle size is odd, then puzzle instance is solvable if number of inversions is even in the input state.
	*/
	/*
		If the puzzle size is even, puzzle instance is solvable if
		the blank is on an odd row counting from the bottom (last, third-last, fifth-last, etc.)
		and number of inversions is even.
	*/
	return inversions%2 == 0
}

func Inversions(puzzle, solvedPuzzle [][]int) (inversions, blankPieceRow int) {
	puzzleSize := len(puzzle) * len(puzzle)
	puzzle1D := make([]int, puzzleSize)
	translate := TranslateMap(solvedPuzzle)
	blankPieceRow = 0
	index := 0
	for i := range puzzle {
		for j := range puzzle[i] {
			if puzzle[i][j] == 0 {
				blankPieceRow = i
			}
			puzzle1D[index] = translate[puzzle[i][j]]
			index++
		}
	}

	zeroPos := ZeroPosition(solvedPuzzle)
	translatedZero := translate[solvedPuzzle[zeroPos.row][zeroPos.column]]
	inversions = 0
	for i := 0; i < len(puzzle1D)-1; i++ {
		for j := i + 1; j < len(puzzle1D); j++ {
			if puzzle1D[i] != translatedZero &&
				puzzle1D[j] != translatedZero && puzzle1D[i] > puzzle1D[j] {
				inversions++
			}
		}
	}

	return inversions, blankPieceRow
}

func TranslateMap(solvedPuzzle [][]int) map[int]int {
	translate := make(map[int]int)

	index := 1
	for i := range solvedPuzzle {
		for j := range solvedPuzzle[i] {
			translate[solvedPuzzle[i][j]] = index
			index++
		}
	}

	return translate
}
