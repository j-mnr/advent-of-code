package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = []byte(`
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)[1:]

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	instrSet := bytes.Split(f, []byte("\n"))
	accum := 0
	i, visited := 0, make(map[int]struct{})
	for _, found := visited[i]; !found; _, found = visited[i] {
		visited[i] = struct{}{}
		switch string(instrSet[i][:3]) {
		case "nop":
			i++
		case "acc":
			n, err := strconv.Atoi(string(instrSet[i][4:]))
			if err != nil {
				log.Fatal(err)
			}
			accum += n
			i++
		case "jmp":
			n, err := strconv.Atoi(string(instrSet[i][4:]))
			if err != nil {
				log.Fatal(err)
			}
			i += n
		}
		fmt.Println(string(instrSet[i]))
	}
	fmt.Println(accum)
}
