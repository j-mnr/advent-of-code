package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var input = []int{
	35,
	20,
	15,
	25,
	47,
	40,
	62,
	55,
	65,
	95,
	102,
	117,
	150,
	182,
	127,
	219,
	299,
	277,
	309,
	576,
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	nums := make([]int, 1000)
	for scr, i := bufio.NewScanner(f), 0; scr.Scan(); i++ {
		n, err := strconv.Atoi(scr.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = n
	}

	preamble := [25]int{}
	sums := map[int]struct{}{}
	for i, n := range nums[:25] {
		preamble[i] = n
	}
	target := 0
	for i, n := range nums[25:] {
		for k := range sums {
			delete(sums, k)
		}
		var j int
		for _, v := range preamble {
			if _, found := sums[n-v]; found {
				break
			} else {
				sums[v] = struct{}{}
			}
			j++
		}
		if j == len(preamble) {
			target = n
			break
		}
		preamble[i%25] = n
	}

	contiguous := make([]int, 0, len(nums))
	for i, n := range nums {
		if n >= target {
			fmt.Println("Something went wrong...")
			break
		}
		contiguous, _ = contiguousSet(nums[i:], target)
		if len(contiguous) > 0 {
			break
		}
	}
	sort.Ints(contiguous)
	fmt.Println(contiguous[len(contiguous)-1] + contiguous[0])
}

func contiguousSet(nums []int, target int) ([]int, error) {
	contiguous := make([]int, 0, len(nums))
	for _, n := range nums {
		target -= n
		contiguous = append(contiguous, n)
		if target == 0 {
			break
		} else if target < 0 {
			return nil, errors.New("These are not the nums you are looking for")
		}
	}
	return contiguous, nil
}
