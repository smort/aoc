package main

import (
	"bufio"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	sc := bufio.NewScanner(strings.NewReader(input))
	if part2 {
		return part2f(sc)
	}
	// solve part 1 here
	return part1(sc)
}

func part1(sc *bufio.Scanner) int {
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		first := -1
		last := -1
		for _, c := range line {
			converted, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}
			if first == -1 {
				first = converted
			}

			last = converted
		}
		sum += (10 * first) + last
	}
	return sum
}

func part2f(sc *bufio.Scanner) int {
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		vals := getNumberValue(line)
		first := -1
		last := -1
		for pos, c := range line {
			converted, err := strconv.Atoi(string(c))
			if err != nil {
				continue
			}
			vals = append(vals, tuple[int, int]{pos, converted})
		}
		slices.SortFunc(vals, func(a, b tuple[int, int]) int {
			if a.One < b.One {
				return -1
			} else if a.One > b.One {
				return 1
			} else {
				return 0
			}
		})

		first = vals[0].Two
		last = vals[len(vals)-1].Two

		sum += (10 * first) + last
	}
	return sum
}

type tuple[T1, T2 any] struct {
	One T1
	Two T2
}

var one = regexp.MustCompile("one")
var two = regexp.MustCompile("two")
var three = regexp.MustCompile("three")
var four = regexp.MustCompile("four")
var five = regexp.MustCompile("five")
var six = regexp.MustCompile("six")
var seven = regexp.MustCompile("seven")
var eight = regexp.MustCompile("eight")
var nine = regexp.MustCompile("nine")

var regexps = []*regexp.Regexp{
	one,
	two,
	three,
	four,
	five,
	six,
	seven,
	eight,
	nine,
}

func getNumberValue(input string) []tuple[int, int] {
	posVal := make([]tuple[int, int], 0)
	for idx, val := range regexps {
		matches := val.FindAllStringIndex(input, -1)
		for _, match := range matches {
			posVal = append(posVal, tuple[int, int]{match[0], idx + 1})
		}
	}

	return posVal
}
