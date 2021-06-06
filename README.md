# n-puzzle

N size puzzle resolver in CLI written in go

You can look at the pdf subject here [n-puzzle subject](https://github.com/vinceduong/n-puzzle/blob/main/npuzzle.en.pdf)

To solve the puzzle i'm using the A* Algorithm to find the shortest path that leads to the solution

## Usage

To create binary file: 
`go install && make build`

Solve puzzle: 
`./n-puzzle  puzzles/3x3-s-1.puzzle`
