package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	passports := make([]map[string]string, 0)
	for _, line := range bytes.Split(f, []byte("\n\n")) {
		p := make(map[string]string)
		for _, sub := range bytes.Fields(line) {
			key, val, _ := bytes.Cut(sub, []byte(":"))
			p[string(key)] = string(val)
		}
		passports = append(passports, p)
	}

	valid := 0
	for _, p := range passports {
		if _, found := p["cid"]; len(p) == 8 || (len(p) == 7 && !found) {
			valid++
		}
	}
	fmt.Println(passports)
	fmt.Println(valid)
}
