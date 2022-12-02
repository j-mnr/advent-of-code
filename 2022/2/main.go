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
	oppRock     choice = "A"
	oppPaper    choice = "B"
	oppScissors choice = "C"

	wantLose choice = "X"
	wantDraw choice = "Y"
	wantWin  choice = "Z"
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
		opp, wantOut, ok := strings.Cut(scr.Text(), " ")
		if !ok {
			fmt.Println("Could not cut", scr.Text())
			continue
		}
		score += outcome(choice(opp), choice(wantOut))
	}
	fmt.Println("score:", score)
}

func outcome(opp, want choice) point {
	switch opp {
	case oppRock:
		switch want {
		case wantLose:
			return lose + scissors
		case wantDraw:
			return draw + rock
		case wantWin:
			return win + paper
		}
	case oppPaper:
		switch want {
		case wantLose:
			return lose + rock
		case wantDraw:
			return draw + paper
		case wantWin:
			return win + scissors
		}
	case oppScissors:
		switch want {
		case wantLose:
			return lose + paper
		case wantDraw:
			return draw + scissors
		case wantWin:
			return win + rock
		}
	}
	panic("Shouldn't reach here")
}
