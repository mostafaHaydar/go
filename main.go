package main

import (
	"fmt"
)

var sudokuArray = [9][9]int{
	{0, 8, 0, 0, 0, 5, 1, 7, 9},
	{0, 0, 0, 2, 0, 6, 0, 8, 4},
	{9, 0, 3, 8, 0, 0, 6, 0, 0},
	{2, 7, 0, 0, 8, 0, 5, 0, 3},
	{4, 0, 0, 0, 5, 0, 8, 1, 2},
	{0, 0, 8, 0, 4, 2, 0, 0, 7},
	{8, 0, 0, 0, 0, 3, 0, 0, 1},
	{3, 5, 4, 0, 1, 0, 0, 9, 0},
	{0, 9, 6, 0, 2, 4, 7, 0, 0},
}

func checkIfTheNumberExistsOnHorizentalLine(lineNumber int, element int) bool {
	var state bool = false
	for i := 0; i < 9; i++ {
		loopElement := sudokuArray[lineNumber][i]
		if loopElement == element {
			state = true
			break
		}
	}
	return state
}

func checkIfTheNumberExistsOnVerticalLine(lineNumber int, element int) bool {
	var state bool = false
	for i := 0; i < 9; i++ {
		loopElement := sudokuArray[i][lineNumber]
		if loopElement == element {
			state = true
			break
		}
	}
	return state
}

func fromWichBlockIsTheElement(positionNumberVertical int,
	positionNumberHorizental int) int {
	if positionNumberVertical < 3 {
		if positionNumberHorizental < 3 {
			return 0
		} else if positionNumberHorizental < 6 {
			return 1
		} else if positionNumberHorizental < 9 {
			return 2
		}
	} else if positionNumberVertical < 6 {
		if positionNumberHorizental < 3 {
			return 3
		} else if positionNumberHorizental < 6 {
			return 4
		} else if positionNumberHorizental < 9 {
			return 5
		}
	} else if positionNumberVertical < 9 {
		if positionNumberHorizental < 3 {
			return 6
		} else if positionNumberHorizental < 6 {
			return 7
		} else if positionNumberHorizental < 9 {
			return 8
		}
	}
	return 0
}
func checkIfExistsInBlock(positionNumberVertical int,
	positionNumberHorizental int, element int) bool {
	blockNumber := fromWichBlockIsTheElement(positionNumberVertical,
		positionNumberHorizental)

	if blockNumber < 3 {
		for i := 0; i < 3; i++ {

			if blockNumber == 0 {
				for j := 0; j < 3; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}
			if blockNumber == 1 {
				for j := 3; j < 6; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}
			if blockNumber == 2 {
				for j := 6; j < 9; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}
		}
	} else if blockNumber < 6 {
		for i := 3; i < 6; i++ {
			if blockNumber == 3 {
				for j := 0; j < 3; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}

			if blockNumber == 4 {
				for j := 3; j < 6; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}

			if blockNumber == 5 {
				for j := 6; j < 9; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}
		}
	} else if blockNumber < 9 {
		for i := 6; i < 9; i++ {

			if blockNumber == 6 {
				for j := 0; j < 3; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}
			if blockNumber == 7 {
				for j := 3; j < 6; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}
			if blockNumber == 8 {
				for j := 6; j < 9; j++ {
					if element == sudokuArray[i][j] {
						return true
					}
				}
			}
		}
	}
	return false
}

func startSolving() {
	var howManyElementsCanBeInOneBox int = 0
	var firstNumberCanBeInTheBox int = 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudokuArray[i][j] == 0 {
				for x := 1; x < 10; x++ {
					if !checkIfExistsInBlock(i, j, x) &&
						!checkIfTheNumberExistsOnVerticalLine(j, x) &&
						!checkIfTheNumberExistsOnHorizentalLine(i, x) {
						firstNumberCanBeInTheBox = x
						howManyElementsCanBeInOneBox = howManyElementsCanBeInOneBox + 1
					}
				}
				if howManyElementsCanBeInOneBox == 1 {
					sudokuArray[i][j] = firstNumberCanBeInTheBox
				}
			}
			firstNumberCanBeInTheBox = 0
			howManyElementsCanBeInOneBox = 0
		}
	}
}

func checkIfexistsZeroInTheSudokuArray() bool {
	state := false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudokuArray[i][j] == 0 {
				state = true
				break
			}
		}
	}
	return state
}

func numberOfEmptyBoxes() int {
	var number int = 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudokuArray[i][j] == 0 {
				number = number + 1
			}
		}
	}
	return number
}

func solvigUntilCompletSolve() {
	var i int = 0
	var emptyBoxesBefore int = 0
	var emptyBoxesAfter int = 0
	for checkIfexistsZeroInTheSudokuArray() {
		emptyBoxesBefore = numberOfEmptyBoxes()
		startSolving()
		emptyBoxesAfter = numberOfEmptyBoxes()

		if emptyBoxesBefore == emptyBoxesAfter {
			break
		}
		i++
	}

	fmt.Printf("-------------------------------------\n")
	for i := 0; i < 9; i++ {
		fmt.Printf("| %d | %d | %d | %d | %d | %d | %d | %d | %d |\n",
			sudokuArray[i][0], sudokuArray[i][1], sudokuArray[i][2],
			sudokuArray[i][3], sudokuArray[i][4], sudokuArray[i][5],
			sudokuArray[i][6], sudokuArray[i][7], sudokuArray[i][8])
		fmt.Printf("-------------------------------------\n")
	}
}

func maindd() {
	solvigUntilCompletSolve()
}
