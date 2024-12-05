package main

import (
	"fmt"
	"strings"
)

const (
	Left          = "left"
	Right         = "right"
	Up            = "up"
	Down          = "down"
	DiagUpLeft    = "diag-up-left"
	DiagUpRight   = "diag-up-right"
	DiagDownLeft  = "diag-down-left"
	DiagDownRight = "diag-down-right"
)

var directions = [8]string{
	Left, Right, Up, Down, DiagUpLeft, DiagUpRight, DiagDownLeft, DiagDownRight,
}

func reverseWord(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func createWordMatrix(m []string) (matrix [][]string) {
	matrix = make([][]string, len(m))

	for i, v := range m {
		matrix[i] = make([]string, len(m))

		for j, k := range strings.Split(v, "") {
			matrix[i][j] = k
		}
	}

	return
}

func outOfBounds(row int, col int, rowLen int, colLen int, w string, direction string) bool {
	wLen := len(w)
	// negative conditions - the ones that are OOB
	left := col-(wLen-1) < 0
	right := col+(wLen-1) >= colLen
	up := row-(wLen-1) < 0
	down := row+(wLen-1) >= rowLen

	switch direction {
	case Left:
		return left
	case Right:
		return right
	case Up:
		return up
	case Down:
		return down
	case DiagUpLeft:
		return up || left
	case DiagUpRight:
		return up || right
	case DiagDownLeft:
		return down || left
	case DiagDownRight:
		return down || right
	default:
		return true
	}
}

func checkWordMatch(matrix [][]string, w string, row int, col int, direction string, reverseMatch bool) bool {
	wLen := len(w)
	var word string

	switch direction {
	case Left:
		for i := col; i > col-wLen; i-- {
			word += matrix[row][i]
		}
	case Right:
		for i := col; i < col+wLen; i++ {
			word += matrix[row][i]
		}
	case Up:
		for i := row; i > row-wLen; i-- {
			word += matrix[i][col]
		}
	case Down:
		for i := row; i < row+wLen; i++ {
			word += matrix[i][col]
		}
	case DiagUpLeft:
		for i := 0; i < wLen; i++ {
			word += matrix[row-i][col-i]
		}
	case DiagUpRight:
		for i := 0; i < wLen; i++ {
			word += matrix[row-i][col+i]
		}
	case DiagDownLeft:
		for i := 0; i < wLen; i++ {
			word += matrix[row+i][col-i]
		}
	case DiagDownRight:
		for i := 0; i < wLen; i++ {
			word += matrix[row+i][col+i]
		}
	}

	if reverseMatch {
		return w == word || w == reverseWord(word)
	} else {
		return w == word
	}
}

// check left -> check i : i - len(word) - 1
// check right -> check i : i - len(word) - 1
// check up -> check i[j] : i[j] - len(word) - 1

func iterateMatrix(m [][]string, w string, dirs []string) int {
	matches := 0
	countByDir := make(map[string]int)

	for i, row := range m {
		for j := range row {
			for _, d := range dirs {
				oob := outOfBounds(i, j, len(m), len(row), w, d)

				if !oob {
					if checkWordMatch(m, w, i, j, d, false) {
						_, ok := countByDir[d]

						if !ok {
							countByDir[d] = 1
						} else {
							countByDir[d]++
						}

						fmt.Printf("dir %s is oob: %v; row: %d, col: %d;\n", d, oob, i, j)
						matches++
						fmt.Printf("matches so far: %d\n", matches)
					}
				}

			}
		}
	}

	fmt.Println(countByDir)

	return matches
}

func iterateMatrixXPattern(m [][]string) int {
	matches := 0
	word := "MAS"

	for i := 1; i < len(m)-1; i++ {
		for j := 1; j < len(m[i])-1; j++ {
			a := m[i-1][j-1] + m[i][j] + m[i+1][j+1]
			b := m[i-1][j+1] + m[i][j] + m[i+1][j-1]

			if (a == word || word == reverseWord(a)) && (b == word || word == reverseWord(b)) {
				matches++
			}

		}
	}

	return matches
}

func SearchForWord(w string, m []string) (occurrences int) {
	matrix := createWordMatrix(m)
	dirs := directions[:]
	occurrences = iterateMatrix(matrix, w, dirs)

	return
}

func SearchForWordXPattern(w string, m []string) (occurrences int) {
	matrix := createWordMatrix(m)
	occurrences = iterateMatrixXPattern(matrix)

	return
}

func main() {

}
