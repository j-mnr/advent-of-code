package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// marker is used to "remove" called out values from the bingo numbers.
	marker = 0
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	bingoNumbers := buildNumbers(strings.Split(s.Text(), ","))
	bingoBoards := makeBoards(s)
	winningBoards := make(map[int]struct{})
	for _, bn := range bingoNumbers {
		for i, board := range bingoBoards {
			for j, row := range board {
				if _, found := winningBoards[i]; found { // Skip winning boards
					break
				}
				for k, n := range row {
					// Mark Bingo Numbers on boards
					if bn == n {
						bingoBoards[i][j][k] = marker
						if checkWin(bingoBoards[i]) {
							winningBoards[i] = struct{}{}
							if len(winningBoards) == len(bingoBoards) {
								fmt.Println(int(bn) * bingoBoards[i].sum())
								return
							}
						}
						goto NEXT_BOARD
					}
				}
			}
		NEXT_BOARD:
		}
	}
}

// checkWin checks board for winning combinations of
//  * entire row full of marker
//  * entire column full of marker
//  * either diagonal full of marker
func checkWin(b bingoBoard) bool {
	rowCnt := 0
	colCnt := [5]int{}
	for _, row := range b {
		for i, n := range row {
			if n == marker {
				rowCnt++
				colCnt[i]++
			}
		}
		if rowCnt == 5 {
			return true
		}
		rowCnt = 0
	}
	for _, c := range colCnt {
		if c == 5 {
			return true
		}
	}
	leftDiag := 0
	rightDiag := 0
	for i := 0; i < len(b); i++ {
		if b[i][i] == marker {
			leftDiag++
		}
		if b[i][len(b)-1-i] == marker {
			rightDiag++
		}
	}
	if leftDiag == 5 || rightDiag == 5 {
		return true
	}
	return false
}

// buildNumbers takes the first line of the input and maps them to uint8s.
func buildNumbers(nums []string) []uint8 {
	bingoNumbers := make([]uint8, len(nums))
	for i, n := range nums {
		bingoNumbers[i] = atoi(n)
	}
	return bingoNumbers
}

// bingoBoard is a 5 by 5 matrix. We can save some space making it uint8
// because the numbers on a bingo board don't go past 255.
type bingoBoard [5][5]uint8

// sum adds all of the numbers remaining on the board.
func (b bingoBoard) sum() int {
	sum := 0
	for _, row := range b {
		for _, n := range row {
			sum += int(n)
		}
	}
	return sum
}

// makeBoards makes all of the 5 by 5 bingo boards from the input.
func makeBoards(s *bufio.Scanner) []bingoBoard {
	bb := make([]bingoBoard, 0, 5)
	n := 0
	row := 0
	for s.Scan() {
		if len(s.Bytes()) == 0 {
			bb = append(bb, bingoBoard{})
			n++
			row = 0
			continue
		}
		for i, num := range strings.Fields(s.Text()) {
			bb[n-1][row][i] = atoi(num)
		}
		row++
	}
	return bb
}

// atoi is a small wrapper to strconv.Atoi to return a uin8. This function will
// panic if the string is not a number.
func atoi(num string) uint8 {
	d, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return uint8(d)
}
