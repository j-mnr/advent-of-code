package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	memRE := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	scr := bufio.NewScanner(f)
	brk := 0
	var mask []byte
	mem := make(map[int]int)
	for scr.Scan() {
		b := scr.Bytes()
		if brk == 4 {
			break
		}
		switch {
		case bytes.HasPrefix(b, []byte("mask")):
			mask = bytes.Fields(scr.Bytes())[2]
			// fmt.Println(string(mask))
		default:
			matches := memRE.FindSubmatch(b)
			addr, err := strconv.Atoi(string(matches[1]))
			if err != nil {
				fmt.Println(err)
			}
			val, err := strconv.Atoi(string(matches[2]))
			val = overwrite(val, mask)
			mem[addr] = val
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println(string(matches[1]), string(matches[2]))
		}
		// brk++
	}
	var sum uint64
	for _, v := range mem {
		sum += uint64(v)
	}
	fmt.Println(sum)
	// fmt.Printf("%#b\n", overwrite(414370178, []byte("01101X001X111X010X0000X1001X010XX0X0")))
	// fmt.Printf("%#b\n", 414370178)
	// 0b01101X001X111X010X0000X1001X010XX0X0
	//
	// 0b011010001011100100000001001101000010
	// 0b000000011000101100101100100110000010
	//
	// 0b011010001011100100000001001101000010
	// 0b000000000000000000000000111101100110
}

func overwrite(n int, mask []byte) int {
	for i, b := range mask {
		switch b {
		case 'X':
			continue
		case '1':
			n |= (1 << (len(mask) - i - 1))
		case '0':
			n &^= (1 << (len(mask) - i - 1))
		}
	}
	return n
}
