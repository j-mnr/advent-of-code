package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var basins []int
	visited := make(map[[2]int]struct{})
	for i, row := range matrix {
		for j, num := range row {
			if _, ok := visited[[2]int{i, j}]; ok || num == 9 {
				continue
			}

			basinSize := 0
			stack := [][2]int{{i, j}}
			for pair := [2]int{}; len(stack) != 0; {
				pair, stack = stack[len(stack)-1], stack[:len(stack)-1]
				i, j := pair[0], pair[1]
				if _, ok := visited[pair]; ok {
					continue
				}
				visited[pair] = struct{}{}
				basinSize++
				stack = append(stack, addNeighbors(i, j, matrix)...)
			}
			basins = append(basins, basinSize)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	fmt.Println(basins[0] * basins[1] * basins[2])
}

func addNeighbors(i, j int, matrix [][]int) [][2]int {
	var nbors [][2]int
	n, m := len(matrix)-1, len(matrix[0])-1
	if i > 0 && matrix[i-1][j] != 9 {
		nbors = append(nbors, [2]int{i - 1, j})
	}
	if j > 0 && matrix[i][j-1] != 9 {
		nbors = append(nbors, [2]int{i, j - 1})
	}
	if i < n && matrix[i+1][j] != 9 {
		nbors = append(nbors, [2]int{i + 1, j})
	}
	if j < m && matrix[i][j+1] != 9 {
		nbors = append(nbors, [2]int{i, j + 1})
	}
	return nbors
}
