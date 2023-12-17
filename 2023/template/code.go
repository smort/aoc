package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	input := readFile("input-example.txt")
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}
