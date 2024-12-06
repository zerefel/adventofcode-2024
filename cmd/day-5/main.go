package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func parseUpdates(u []string) (updates [][]int) {
	updates = make([][]int, len(u))

	for i, line := range u {
		update := strings.Split(line, ",")

		updates[i] = make([]int, len(update))

		for j, v := range update {
			num, _ := strconv.Atoi(v)
			updates[i][j] = num
		}

	}

	return
}

func isUpdateCorrect(update []int, rules []map[int]int) bool {
	for i, num := range update {
		for _, rule := range rules {
			for left, right := range rule {
				leftIndex := slices.Index(update, left)
				rightIndex := slices.Index(update, right)

				if num == left && rightIndex != -1 && rightIndex < i {
					return false
				} else if num == right && leftIndex != -1 && slices.Index(update, left) > i {
					return false
				}
			}
		}
	}

	return true
}

func findMidpoint(update []int) int {
	return update[int(math.Floor(float64(len(update)/2)))]
}

func fixIncorrectUpdate(update []int, rules []map[int]int) []int {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(update); i++ {
			// num := update[i]

			for _, rule := range rules {
				for left, right := range rule {
					leftIndex := slices.Index(update, left)
					rightIndex := slices.Index(update, right)

					if leftIndex == -1 || rightIndex == -1 {
						continue
					}

					// If 'left' comes after 'right', swap them
					if leftIndex > rightIndex {
						// Perform the swap
						update[leftIndex], update[rightIndex] = update[rightIndex], update[leftIndex]
						swapped = true
					}
				}
			}
		}
	}

	return update
}

func FindCorrectlyOrderedUpdates(rules []string, updates []string) (correctUpdates int) {
	parsedUpdates := parseUpdates(updates)
	ruleMapArr := make([]map[int]int, 0)
	totalCorrectCount := 0

	for _, rule := range rules {
		ruleMap := make(map[int]int)
		ruleParts := strings.Split(rule, "|")

		left, _ := strconv.Atoi(ruleParts[0])
		right, _ := strconv.Atoi(ruleParts[1])
		ruleMap[left] = right

		ruleMapArr = append(ruleMapArr, ruleMap)
	}

	for _, update := range parsedUpdates {
		if isUpdateCorrect(update, ruleMapArr) {
			correctUpdates += findMidpoint(update)
			totalCorrectCount++
		}
	}

	fmt.Printf("total correct updates count %d ", totalCorrectCount)

	return correctUpdates
}

func FixIncorrectlyOrderedUpdates(rules []string, updates []string) (fixedUpdates int) {
	parsedUpdates := parseUpdates(updates)
	ruleMapArr := make([]map[int]int, 0)
	totalIncorrect := 0
	totalFixed := 0

	for _, rule := range rules {
		ruleMap := make(map[int]int)
		ruleParts := strings.Split(rule, "|")

		left, _ := strconv.Atoi(ruleParts[0])
		right, _ := strconv.Atoi(ruleParts[1])
		ruleMap[left] = right

		ruleMapArr = append(ruleMapArr, ruleMap)
	}

	for _, update := range parsedUpdates {
		if !isUpdateCorrect(update, ruleMapArr) {
			totalIncorrect++
			update = fixIncorrectUpdate(update, ruleMapArr)

			if isUpdateCorrect(update, ruleMapArr) {
				totalFixed++
				fixedUpdates += findMidpoint(update)
			}
		}
	}

	fmt.Printf("Total incorrect count: %d", totalIncorrect)
	fmt.Printf("Total fixed count: %d", totalFixed)

	return fixedUpdates
}
