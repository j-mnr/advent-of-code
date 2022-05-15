package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// TODO(jaym): part two
func main() {
	f, _ := os.Open("input")
	defer f.Close()
	epsilon, gamma := buildMostAndLeastCommonNum(countBits(f))
	fmt.Println(epsilon * gamma)
}

func buildMostAndLeastCommonNum(cnt []oneCount, lnCnt int) (x, y int) {
	for _, oc := range cnt {
		x <<= 1
		y <<= 1
		if int(oc) > lnCnt/2 {
			x++
		} else {
			y++
		}
	}
	return x, y
}

type oneCount uint64

func countBits(r io.Reader) ([]oneCount, int) {
	scn := bufio.NewScanner(r)
	scn.Scan()
	cnt := make([]oneCount, len(scn.Text()))
	addCount(cnt, scn.Text())
	lnCount := 1
	for scn.Scan() {
		lnCount++
		addCount(cnt, scn.Text())
	}
	return cnt, lnCount
}

func addCount(c []oneCount, bits string) {
	for i, r := range bits {
		if r == '1' {
			c[i]++
		}
	}
}
