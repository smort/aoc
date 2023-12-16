package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type sourceDestMap struct {
	Source      int
	Destination int
	Length      int
}

type maps []sourceDestMap

func main() {
	input := readFile("input-example.txt")
	seeds, maps := parseInput(input)
	fmt.Println(seeds, maps)
}

// first return is seeds, second is maps
func parseInput(input []string) ([]int, []maps) {
	seedSplit := strings.Split(input[0], "seeds: ")
	seedStrs := strings.Split(strings.TrimSpace(seedSplit[1]), " ")
	seeds := make([]int, 0, len(seedStrs))
	for _, seed := range seedStrs {
		parsed, err := strconv.ParseInt(seed, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, int(parsed))
	}

	m := make([]maps, 0, 7)
	var currThings maps

	i := 1
	for i < len(input) {
		line := input[i]
		if strings.TrimSpace(line) == "" || strings.Contains(line, ":") {
			if currThings != nil {
				m = append(m, currThings)
				currThings = nil
			}
			i++
			continue
		}
		if currThings == nil {
			currThings = make(maps, 0)
		}

		vals := strings.Split(line, " ")
		thing := sourceDestMap{}
		for i, val := range vals {
			valInt, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			if i == 0 {
				thing.Destination = int(valInt)
			} else if i == 1 {
				thing.Source = int(valInt)
			} else {
				thing.Length = int(valInt)
			}
		}
		currThings = append(currThings, thing)
		i++
	}

	return seeds, m
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}
