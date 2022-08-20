package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var r = strings.NewReader(`
2199943210
3987894921
9856789892
8767896789
9899965678`[1:])

func main() {
	r, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scr := bufio.NewScanner(r)
	var matrix [][]int
	for scr.Scan() {
		var nums []int
		for _, d := range scr.Text() {
			nums = append(nums, int(d-'0'))
		}
		matrix = append(matrix, nums)
	}

	lowsSum := 0
	for i, row := range matrix {
		for j, num := range row {
			if isLowPoint(num, i, j, matrix) {
				fmt.Println(num, i, j)
				lowsSum += num + 1
			}
		}
	}
	fmt.Println(lowsSum)
}

func isLowPoint(num, i, j int, matrix [][]int) bool {
	n, m := len(matrix)-1, len(matrix[0])-1
	if (i > 0 && num >= matrix[i-1][j]) || (j > 0 && num >= matrix[i][j-1]) ||
		(i < n && num >= matrix[i+1][j]) || (j < m && num >= matrix[i][j+1]) {
		return false
	}
	return true
}
