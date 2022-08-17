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

const (
	target = 19690720
)

type instruction struct {
	Op         Opcode
	Noun, Verb uint
	OutAddress uint
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	var instructions []instruction
	for i, sepInstructions := 0, bytes.Split(f, []byte(",")); i < len(sepInstructions); i += 4 {
		op := Opcode(atoui(sepInstructions[i]))
		if op == 99 {
			break
		}
		instructions = append(instructions, instruction{
			Op:         op,
			Noun:       atoui(sepInstructions[i+1]),
			Verb:       atoui(sepInstructions[i+2]),
			OutAddress: atoui(sepInstructions[i+3]),
		})
	}
	cpyInstr := make([]instruction, len(instructions))
	copy(cpyInstr, instructions)
	var got uint
	var lo, hi uint = 0, 99
	for got != target && hi > 0 && lo < 99 {
		copy(cpyInstr, instructions)
		cpyInstr[0].Noun = lo
		cpyInstr[0].Verb = hi
		for _, instr := range cpyInstr {
			rVal1, rVal2 := load(cpyInstr, instr.Noun), load(cpyInstr, instr.Verb)
			outReg := load(cpyInstr, instr.OutAddress)
			switch instr.Op {
			case 1:
				*outReg = *rVal1 + *rVal2
			case 2:
				*outReg = *rVal1 * *rVal2
			}
		}
		got = uint(cpyInstr[0].Op)
		switch {
		case got > target:
			hi--
		case got < target:
			lo++
		}
	}
	fmt.Println(100*cpyInstr[0].Noun + cpyInstr[0].Verb)
}

func load(instructions []instruction, registerPos uint) *uint {
	reg := registerPos % 4
	switch reg {
	case 0:
		return ((*uint)(unsafe.Pointer(&instructions[registerPos/4].Op)))
	case 1:
		return &instructions[registerPos/4].Noun
	case 2:
		return &instructions[registerPos/4].Verb
	case 3:
		return &instructions[registerPos/4].OutAddress
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
