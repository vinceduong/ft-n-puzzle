package solve

import "fmt"

func PuzzleToString(p [][]int) string {
	puzzleString := ""
	for i := range p {
		for j := range p[i] {
			if (i == len(p) - 1 && j == len(p) - 1) {
				puzzleString += fmt.Sprintf("%v", p[i][j]) 
			} else {
				puzzleString += fmt.Sprintf("%v-", p[i][j]) 
			}
		}
	}

	return puzzleString
}
