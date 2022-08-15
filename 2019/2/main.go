package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

type Opcode uint

const (
	Add Opcode = iota + 1
	Multiply

	Halt Opcode = 99
)

type instruction struct {
	Op                       Opcode
	InRegister1, InRegister2 uint
	OutRegister              uint
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	var instructions []instruction
	sepInstructions := bytes.Split(f, []byte(","))
	i := 0
	for ; i < len(sepInstructions); i += 4 {
		op := Opcode(atoui(sepInstructions[i]))
		if op == 99 {
			break
		}
		instructions = append(instructions, instruction{
			Op:          op,
			InRegister1: atoui(sepInstructions[i+1]),
			InRegister2: atoui(sepInstructions[i+2]),
			OutRegister: atoui(sepInstructions[i+3]),
		})
	}
	// Part 1 shenanigans
	instructions[0].InRegister1 = 12
	instructions[0].InRegister2 = 2
	// NOTE(jay): This may be needed for Part 2?
	// var finalRegs []uint
	// for ; i < len(sepInstructions); i++ {
	// 	finalRegs = append(finalRegs, atoui(bytes.TrimSpace(sepInstructions[i])))
	// }
	for _, instr := range instructions {
		rVal1, rVal2 := load(instructions, instr.InRegister1), load(instructions, instr.InRegister2)
		outReg := load(instructions, instr.OutRegister)
		switch instr.Op {
		case 1:
			*outReg = *rVal1 + *rVal2
		case 2:
			*outReg = *rVal1 * *rVal2
		}
	}
	fmt.Println(instructions[0].Op)
}

func load(instructions []instruction, registerPos uint) *uint {
	reg := registerPos % 4
	switch reg {
	case 0:
		return ((*uint)(unsafe.Pointer(&instructions[registerPos/4].Op)))
	case 1:
		return &instructions[registerPos/4].InRegister1
	case 2:
		return &instructions[registerPos/4].InRegister2
	case 3:
		return &instructions[registerPos/4].OutRegister
	default:
		panic("This should never be reached")
	}
}

func atoui(b []byte) uint {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatal(err)
	}
	return uint(n)
}
