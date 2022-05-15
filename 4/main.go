package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	marker = 0
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	bingoNumbers := buildNumbers(strings.Split(s.Text(), ","))
	// Make Boards
	fmt.Println(bingoNumbers)
	bingoBoards := makeBoards(s)
	for _, bn := range bingoNumbers {
		for i, board := range bingoBoards {
			for j, row := range board {
				for k, n := range row {
					// Mark Bingo Numbers on boards
					if bn == n {
						bingoBoards[i][j][k] = marker
						// Check board for winner -- needs logic
						//  * Check rows all have marker
						//  * Check cols all have marker
						//  * Check diags all have marker
						if checkWin(bingoBoards[i]) {
							// Calculate Winning board score
							fmt.Println(int(bn) * sum(bingoBoards[i]))
							goto END
						}
						goto NEXT_BOARD
					}
				}
			}
		NEXT_BOARD:
		}
	}
END:
	// for _, bb := range bingoBoards {
	// 	fmt.Println(bb)
	// }
}

func sum(b bingoBoard) int {
	sum := 0
	for _, row := range b {
		for _, n := range row {
			sum += int(n)
		}
	}
	return sum
}

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
			fmt.Println(b)
			return true
		}
		rowCnt = 0
	}
	for _, c := range colCnt {
		if c == 5 {
			fmt.Println(b)
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
		fmt.Println(b)
		return true
	}
	return false
}

func buildNumbers(nums []string) []uint8 {
	bingoNumbers := make([]uint8, len(nums))
	for i, n := range nums {
		bingoNumbers[i] = atoi(n)
	}
	return bingoNumbers
}

type bingoBoard [5][5]uint8

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
	// for _, bb := range bb {
	// 	fmt.Println(bb)
	// }
	return bb
}

func atoi(num string) uint8 {
	d, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return uint8(d)
}
