package main

import (
	"log"
	"os"

	"github.com/vinceduong/n-puzzle/src/parse"
	"github.com/vinceduong/n-puzzle/src/solve"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No file provided")
	}
	fileName := os.Args[1]

	lines := parse.LinesFromFile(fileName)
	puzzle, _ := parse.PuzzleFromLines(lines)

	solve.Astar(puzzle)
}
