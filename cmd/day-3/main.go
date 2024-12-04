package main

import (
	"regexp"
	"strconv"
)

func extractEnabledOps(s string) (ops string) {
	r := regexp.MustCompile(`(?:^(.*?)don't\(\))|(?:do\(\)(.*?)(?:don't\(\)|$))`)
	matches := r.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		if match[1] != "" {
			ops += match[1]
		}
		if match[2] != "" {
			ops += match[2]
		}
	}

	return
}

func parseCorruptedMulOps(s string) (ops [][]int) {
	r := regexp.MustCompile(`mul\((?P<a>\d{1,3})\,(?P<b>\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(s, -1)
	ops = make([][]int, len(matches))

	for i, m := range matches {
		ops[i] = make([]int, 2)

		if len(m) > 2 {
			a, err1 := strconv.Atoi(m[1])
			b, err2 := strconv.Atoi(m[2])

			if err1 == nil || err2 == nil {
				ops[i][0] = a
				ops[i][1] = b
			}

		}
	}

	return
}

func calculateSum(ops [][]int) (sum int) {
	for _, nums := range ops {
		mul := 1

		for _, n := range nums {
			mul *= n
		}

		sum += mul
	}

	return
}

func CalculateCorruptOpsSum(str string, extractOps bool) (sum int) {
	if extractOps {
		str = extractEnabledOps(str)
	}

	ops := parseCorruptedMulOps(str)
	sum = calculateSum(ops)

	return
}

func main() {

}
