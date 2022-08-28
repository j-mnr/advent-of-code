package main

import "fmt"

// state is the previous and current turn values of a number in the map
type state struct{ prev, curr int }

func main() {
	input := []int{13, 16, 0, 12, 15, 1}
	m := map[int]state{}
	for i, n := range input {
		m[n] = state{curr: i + 1}
	}
	prevNum := input[len(input)-1]
	// TODO(jay): Take this back down to 2020 for part 1
	for turn := len(input) + 1; turn <= 30000000; turn++ {
		s, _ := m[prevNum]
		if s.prev == 0 {
			prevNum = s.prev
		} else {
			prevNum = s.curr - s.prev
		}
		m[prevNum] = state{prev: m[prevNum].curr, curr: turn}
		if turn == 30000000 {
			fmt.Println(prevNum)
		}
	}
}
