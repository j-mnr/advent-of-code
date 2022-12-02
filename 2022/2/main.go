package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type choice string

const (
	opRock     choice = "A"
	opPaper    choice = "B"
	opScissors choice = "C"

	myRock     choice = "X"
	myPaper    choice = "Y"
	myScissors choice = "Z"
)

type point uint

const (
	lose point = iota * 3
	draw
	win
)

const (
	rock point = iota + 1
	paper
	scissors
)

var f = strings.NewReader(`
A Y
B X
C Z`[1:])

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	score := point(0)
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		op, my, ok := strings.Cut(scr.Text(), " ")
		if !ok {
			fmt.Println("Could not cut", scr.Text())
			continue
		}
		score += outcome(choice(op), choice(my))
		score += convert(choice(my))
	}
	fmt.Println("score:", score)
}

func outcome(op, my choice) point {
	switch op {
	case opRock:
		switch my {
		case myRock:
			return draw
		case myPaper:
			return win
		case myScissors:
			return lose
		}
	case opPaper:
		switch my {
		case myRock:
			return lose
		case myPaper:
			return draw
		case myScissors:
			return win
		}
	case opScissors:
		switch my {
		case myRock:
			return win
		case myPaper:
			return lose
		case myScissors:
			return draw
		}
	}
	panic("Shouldn't reach here")
}

func convert(my choice) point {
	switch my {
	case myScissors:
		return scissors
	case myPaper:
		return paper
	case myRock:
		return rock
	default:
		panic("Shouldn't reach here")
	}
}
