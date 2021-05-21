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
        return Position{row: i, column: j}
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

func Neighbors(n *Node) []*Node {
  puzzleSize := len(n.puzzle)
  potentialZeroPositions := PotentialZeroPositions(n.zeroPosition, puzzleSize)
  neighbors := make([]*Node, len(potentialZeroPositions))

  for i, position := range potentialZeroPositions {
    newPuzzle := SwapPuzzlePieces(n.puzzle, n.zeroPosition, position)
    cost := n.cost + 1

    neighbors[i] = &Node{
      puzzle: newPuzzle,
      cost: cost,
      heuristic: 0,
      score: 0,
      zeroPosition: position,
      parent: n,
    }
  }

  return neighbors
}
