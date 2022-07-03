package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var validationFuncs = map[string]func(string) bool{
	"byr": func(v string) bool {
		n, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		return 1920 <= n && n <= 2002
	},
	"iyr": func(v string) bool {
		n, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		return 2010 <= n && n <= 2020
	},
	"eyr": func(v string) bool {
		n, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		return 2020 <= n && n <= 2030
	},
	"hgt": func(v string) bool {
		matches := regexp.MustCompile(`^(\d+)(cm|in)$`).FindStringSubmatch(v)
		if len(matches) < 3 {
			return false
		}
		n, err := strconv.Atoi(matches[1])
		if err != nil {
			return false
		}
		switch matches[2] {
		case "cm":
			return 150 <= n && n <= 193
		case "in":
			return 59 <= n && n <= 76
		default:
			return false
		}
	},
	"hcl": func(v string) bool {
		return regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(v)
	},
	"ecl": func(v string) bool {
		switch v {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		default:
			return false
		}
	},
	"pid": func(v string) bool {
		return regexp.MustCompile(`^[0-9]{9}$`).MatchString(v)
	},
	"cid": func(v string) bool { return true },
}

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
		if isValid(p) {
			valid++
		}
	}
	fmt.Println(valid)
}

func isValid(p map[string]string) bool {
	for k, v := range p {
		if !validationFuncs[k](v) {
			return false
		}
	}
	if _, found := p["cid"]; len(p) == 8 || (len(p) == 7 && !found) {
		return true
	}
	return false
}
