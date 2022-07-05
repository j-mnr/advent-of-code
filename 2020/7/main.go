package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
)

type bag struct {
	name     string
	children []*bag
	parents  []*bag
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	bags := make(map[string]*bag)
	for scr := bufio.NewScanner(f); scr.Scan(); {
		pName, kids, found := bytes.Cut(scr.Bytes(), []byte(" bags contain "))
		if !found {
			log.Fatal(pName, kids, "No rule here")
		}
		parent := &bag{name: string(pName)}
		for _, c := range bytes.Split(kids, []byte(",")) {
			childRe := regexp.MustCompile(`(\d+)\s([a-z\s]+)\sbags?.?`)
			matches := childRe.FindSubmatch(c)
			if len(matches) < 3 {
				for _, m := range matches {
					log.Println(m)
				}
				continue
			}
			parent.children = append(parent.children, &bag{name: string(matches[2])})
		}
		bags[parent.name] = parent
	}
	// add the parents to the nodes
	for pName := range bags {
		for i := range bags[pName].children {
			child, ok := bags[bags[pName].children[i].name]
			if !ok {
				log.Println("not found")
				continue
			}
			child.parents = append(child.parents, bags[pName])
			bags[pName].children[i] = child
		}
	}
	// dfs
	start := bags["shiny gold"]
	next := []*bag{start}
	var n *bag
	visited := make(map[string]struct{})
	for len(next) > 0 {
		n, next = next[0], next[1:]
		for _, p := range n.parents {
			next = append(next, p)
		}
		visited[n.name] = struct{}{}
	}
	fmt.Println(len(visited) - 1)
}
