package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	maxSize        = 100000
	availableSpace = 70000000
	needSpace      = 30000000
)

type fileType uint8

const (
	unknown fileType = iota
	dir
	file
)

func (ft fileType) String() string {
	switch ft {
	case unknown:
		return "unknown"
	case dir:
		return "dir"
	case file:
		return "file"
	default:
		return "NA"
	}
}

type resource struct {
	name     string
	size     uint64
	typ      fileType
	parent   *resource
	children []*resource
}

func (r *resource) Size() uint64 {
	if r.size != 0 {
		return r.size
	}
	sum := uint64(0)
	for _, c := range r.children {
		switch c.typ {
		case file:
			sum += c.size
		case dir:
			sum += c.Size()
		}
	}
	r.size = sum
	return r.size
}

func (r *resource) Directories() []*resource {
	var dirs []*resource
	for _, c := range r.children {
		switch c.typ {
		case file:
			// nop
		case dir:
			dirs = append(dirs, c)
			dirs = append(dirs, c.Directories()...)
		}
	}
	return dirs
}

func (r resource) String() string {
	var sb strings.Builder
	spacing := 0
	for n := &r; n.parent != nil; n = n.parent {
		spacing += 2
	}
	sb.WriteString(fmt.Sprintf("%s- %-4s % 8d %s\n",
		strings.Repeat(" ", spacing),
		r.typ, r.size, r.name,
	))
	for _, c := range r.children {
		sb.WriteString(c.String())
	}
	return sb.String()
}

var f = strings.NewReader(`
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`[1:])

func main() {
	root := &resource{typ: dir, children: make([]*resource, 0), name: "/"}
	node := root
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(data, []byte("\n"))
	lines = lines[:len(lines)-1]
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		switch { // commands
		case bytes.HasPrefix(line, []byte("$ cd")):
			if bytes.Equal(line, []byte("$ cd /")) {
				continue
			}
			if bytes.Equal(line, []byte("$ cd ..")) {
				node = node.parent
			}

			name := bytes.Fields(line)[2]
			for _, c := range node.children {
				if c.name == string(name) {
					node = c
				}
			}
		case bytes.HasPrefix(line, []byte("$ ls")):
			i++
			for i < len(lines) && !bytes.Contains(lines[i], []byte("$")) {
				r := parse(lines[i])
				r.parent = node
				node.children = append(node.children, r)
				i++
			}
			i--
		default:
			panic("Other commands not supported" + string(line))
		}
	}
	root.Size()
	total := uint64(0)
	for _, d := range root.Directories() {
		if d.size >= maxSize {
			continue
		}
		total += d.size
	}

	// Part 2
	usedSpace := availableSpace - root.size
	minSpace := uint64(math.MaxUint64)
	minDirSize := uint64(0)
	for _, d := range root.Directories() {
		if addedSpace := d.size + usedSpace; addedSpace < minSpace && addedSpace >= needSpace {
			minSpace = addedSpace
			minDirSize = d.size
		}
	}
	fmt.Printf("minDirSize: %v\n", minDirSize)
}

func parse(line []byte) *resource {
	f := bytes.Fields(line)
	switch {
	case bytes.HasPrefix(f[0], []byte("dir")):
		return &resource{
			name: string(f[1]),
			typ:  dir,
		}
	default: // Must be a file
		n, err := strconv.Atoi(string(f[0]))
		if err != nil {
			log.Fatal(err)
		}
		return &resource{
			name: string(f[1]),
			size: uint64(n),
			typ:  file,
		}
	}
}
