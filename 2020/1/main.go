package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var input = strings.NewReader(`
1721
979
366
299
675
1456`[1:])

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	nums := Nums(f)
	fmt.Println(nums)
	for i, j := 0, len(nums)-1; i <= j; {
		sum := nums[i] + nums[j]
		switch {
		case sum < 2020:
			i++
		case sum > 2020:
			j--
		case sum == 2020:
			fmt.Println(nums[i], nums[j], nums[i]*nums[j])
			return
		}
	}
}

func Nums(r io.Reader) []int {
	nums := make([]int, 0, 1000)
	scr := bufio.NewScanner(r)
	for scr.Scan() {
		n, err := strconv.Atoi(scr.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	sort.Ints(nums)
	return nums
}
