package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vinceduong/n-puzzle/src/parse"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No file provided")
	}
	fileName := os.Args[1]

	lines := parse.GetLinesFromFile(fileName)
	fmt.Println("'" + strings.Join(lines, `','`) + `'`)
	puzzle, puzzleSize := parse.GetPuzzleFromLines(lines)
	fmt.Println(puzzle)
	fmt.Println(puzzleSize)
}
