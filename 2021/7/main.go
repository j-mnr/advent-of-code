package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	// r := strings.NewReader("16,1,2,0,4,2,7,1,2,14")
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	nums := setup(f)

	consumption, last := make([][]int, max(nums)), len(nums)
	for i := range consumption {
		consumption[i] = make([]int, last)
	}
	minFuel := (1 << 63) - 1
	for i, row := range consumption {
		for j := range row {
			n := nums[j]
			consumption[i][j] = step2(int((math.Abs(float64(n - i)))))
		}
		if m := sum(consumption[i]); m < minFuel {
			minFuel = m
		}
	}

	fmt.Println(minFuel)
}

func step2(n int) (s int) {
	for i := 1; i <= n; i++ {
		s += i
	}
	return s
}

func sum(nums []int) (s int) {
	for _, n := range nums {
		s += n
	}
	return s
}

func max(nums []int) (m int) {
	for _, n := range nums {
		if m < n {
			m = n
		}
	}
	return m
}

func setup(r io.Reader) (nums []int) {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	for _, d := range bytes.Split((b), []byte{','}) {
		d = bytes.TrimSpace(d)
		n, err := strconv.Atoi(string(d))
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}
