package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLinesFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		lines    []string
		bytes    []byte
		isPrefix bool = false
	)

	reader := bufio.NewReader(file)
	for !isPrefix && err == nil {
		bytes, isPrefix, err = reader.ReadLine()
		lines = append(lines, string(bytes))
	}

	return lines
}

func getPuzzleFromLines(lines []string) ([][]int, int) {
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

func main() {
	fileName := os.Args[1]

	lines := getLinesFromFile(fileName)
	fmt.Println("'" + strings.Join(lines, `','`) + `'`)
	puzzle, puzzleSize := getPuzzleFromLines(lines)
	fmt.Println(puzzle)
	fmt.Println(puzzleSize)
}
