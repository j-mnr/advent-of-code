package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var octopi [10][10]int
	scr := bufio.NewScanner(r)
	for i := 0; scr.Scan(); i++ {
		for j, b := range scr.Text() {
			octopi[i][j] = int(b - '0')
		}
	}

	nSteps := 0
	for !allFlashed(octopi) {
		nSteps++
		var queue [][2]int
		for i, row := range octopi {
			for j := range row {
				octopi[i][j]++
				if octopi[i][j] < 10 {
					continue
				}
				octopi[i][j] = 0
				queue = append(queue, getNeighbors(i, j, octopi)...)
			}
		}

		for pair := [2]int{}; len(queue) != 0; {
			pair, queue = queue[0], queue[1:]
			i, j := pair[0], pair[1]
			if octopi[i][j] == 0 { // Don't give more energy
				continue
			}
			octopi[i][j]++
			if octopi[i][j] < 10 {
				continue
			}
			octopi[i][j] = 0
			queue = append(queue, getNeighbors(i, j, octopi)...)
		}
	}

	fmt.Println(nSteps)
}

func getNeighbors(i, j int, octopi [10][10]int) (queue [][2]int) {
	if i > 0 {
		queue = append(queue, [2]int{i - 1, j})
	}
	if j > 0 {
		queue = append(queue, [2]int{i, j - 1})
	}
	if i < len(octopi)-1 {
		queue = append(queue, [2]int{i + 1, j})
	}
	if j < len(octopi[0])-1 {
		queue = append(queue, [2]int{i, j + 1})
	}
	if i > 0 && j > 0 {
		queue = append(queue, [2]int{i - 1, j - 1})
	}
	if i < len(octopi)-1 && j < len(octopi[0])-1 {
		queue = append(queue, [2]int{i + 1, j + 1})
	}
	if i > 0 && j < len(octopi[0])-1 {
		queue = append(queue, [2]int{i - 1, j + 1})
	}
	if j > 0 && i < len(octopi)-1 {
		queue = append(queue, [2]int{i + 1, j - 1})
	}
	return queue
}

func allFlashed(sl [10][10]int) bool {
	for _, row := range sl {
		for _, n := range row {
			if n != 0 {
				return false
			}
		}
	}
	return true
}
