package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	tsAndIDs := bytes.Split(f, []byte("\n"))
	timestamp, err := strconv.Atoi(string(tsAndIDs[0]))
	if err != nil {
		log.Fatal(err)
	}

	minID, diff := math.MaxInt, math.MaxInt
	for _, b := range bytes.FieldsFunc(tsAndIDs[1], func(r rune) bool { return r == ',' || r == 'x' }) {
		n, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatal(err)
		}
		if dif := ((timestamp/n)+1)*(n) - timestamp; dif < diff {
			minID = n
			diff = dif
		}
	}
	fmt.Println(minID, diff, minID*diff)
}
