package main

import (
	"fmt"
	"slices"
	"strings"
)

const (
	Up    = "up"
	Right = "right"
	Down  = "down"
	Left  = "left"
)

var directionsOrder = []string{
	Up, Right, Down, Left,
}

func getNextDirection(cur string) string {
	index := slices.Index(directionsOrder, cur)

	if index == len(directionsOrder)-1 {
		index = 0
	} else {
		index++
	}

	return directionsOrder[index]
}

func isObstacle(elem string) bool {
	return elem == "#"
}

func isOOB(row int, col int, labMap [][]string) bool {
	rowLen := len(labMap)
	colLen := len(labMap[0])

	return row < 0 || row > rowLen-1 || col < 0 || col > colLen-1
}

func walkLabMap(labMap [][]string, row int, col int, direction string, limit int) (positions int) {
	for {
		nextRow := row
		nextCol := col

		switch direction {
		case Up:
			nextRow--
		case Right:
			nextCol++
		case Down:
			nextRow++
		case Left:
			nextCol--
		}

		if isOOB(nextRow, nextCol, labMap) {
			break
		}

		if positions > limit {
			fmt.Println("Limit reached, perhaps an infinite loop...")
			return -1
		}
		elem := labMap[nextRow][nextCol]

		if isObstacle(elem) {
			direction = getNextDirection(direction)
			positions++
		} else {
			row = nextRow
			col = nextCol

			if labMap[row][col] == "." {
				labMap[row][col] = "X"
				positions++
			}
		}

	}

	return
}

func findStartPoint(labMap [][]string) (int, int) {
	for i, row := range labMap {
		for j, col := range row {
			if col == "^" {
				return i, j
			}
		}
	}

	return 0, 0
}

func parseLabMap(rawMap []string) (labMap [][]string) {
	labMap = make([][]string, len(rawMap))

	for i, v := range rawMap {
		entries := strings.Split(v, "")
		labMap[i] = entries
	}

	return
}

func ExitLab(rawMap []string) (positions int) {
	limit := 100000
	labMap := parseLabMap(rawMap)
	row, col := findStartPoint(labMap)
	count := walkLabMap(labMap, row, col, Up, limit) + 1

	fmt.Println(count)

	return
}

func ExitLabWithObstruction(rawMap []string) (positions int) {

	labMap := parseLabMap(rawMap)
	startRow, startCol := findStartPoint(labMap)
	obstructionsCount := 0
	prevIndices := [2]int{}

	for i, row := range labMap {
		for j := range row {
			if isObstacle(labMap[i][j]) {
				continue
			}

			labMap[prevIndices[0]][prevIndices[1]] = "."

			labMap[i][j] = "#"
			prevIndices[0] = i
			prevIndices[1] = j
			limit := 100000

			count := walkLabMap(labMap, startRow, startCol, Up, limit)

			if count == -1 {
				obstructionsCount++
			}
		}
	}

	fmt.Println(obstructionsCount)

	return
}

func main() {

	var sampleLab = []string{}

	// var sampleLab = []string{
	// 	"....#.....",
	// 	".........#",
	// 	"..........",
	// 	"..#.......",
	// 	".......#..",
	// 	"..........",
	// 	".#..^.....",
	// 	"........#.",
	// 	"#.........",
	// 	"......#...",
	// }

	ExitLabWithObstruction(sampleLab)
}
