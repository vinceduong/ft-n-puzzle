package parse

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetPuzzleFromLines(lines []string) ([][]int, int) {
	var (
		err        error
		currentRow int = 0
		puzzleSize int = -1
		puzzle     [][]int
	)

	for _, line := range lines {
		fmt.Printf("\nLine: \"%v\"\n", line)
		if len(line) == 0 || line[0] == '#' {
			fmt.Println("Comment")
			continue
		}

		if puzzleSize == -1 {
			puzzleSize, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Puzzle size: %v\n", puzzleSize)

			puzzle = make([][]int, puzzleSize)
			for i := range puzzle {
				puzzle[i] = make([]int, puzzleSize)
			}

			continue
		}

		fmt.Printf("Current Row: %d\n", currentRow)
		numbers := strings.Split(line, " ")
		for i, stringNumber := range numbers {

			if i == puzzleSize {
				log.Fatal("File format is not valid")
			}

			number, err := strconv.Atoi(stringNumber)
			if err != nil {
				log.Fatal(err)
			}

			puzzle[currentRow][i] = number
		}
		currentRow++
	}

	return puzzle, puzzleSize
}
