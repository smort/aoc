package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	input := readFile("input-user.txt")
	// input := readFile("input-example.txt")
	// input := readFile("input-example-2.txt")
	input = input[:len(input)-1]

	gridSlice := make([][]string, 0, len(input))
	for _, line := range input {
		chars := strings.Split(line, "")
		gridSlice = append(gridSlice, chars)
	}

	g := New(gridSlice)

	fmt.Println("Part 1:", part1F(g))
	fmt.Println("Part 2:", part2F(g))
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}

func part1F(g grid) int {
	moves := 0
	for {
		g.MoveToNext()
		moves++
		if g.CurrChar() == "S" {
			break
		}
	}
	return moves / 2
}

func part2F(g grid) int {
	moves := 0
	verticies := make([][2]int, 0)

	if g.IsSVertex() {
		g.GoToStart()
		verticies = append(verticies, g.Pos)
	}

	g.GoToStart() // awkward

	for {
		g.MoveToNext()
		currChar := g.CharAt(g.Pos[0], g.Pos[1])
		if IsVertex(currChar) {
			verticies = append(verticies, g.Pos)
		}

		moves++
		if g.CurrChar() == "S" {
			break
		}
	}

	area := math.Abs(float64(Shoelace(verticies)))
	return PicksTheorem(int(area), moves)
}

var possibleDirections = map[string][2]int{
	"up":    {0, -1},
	"right": {1, 0},
	"down":  {0, 1},
	"left":  {-1, 0},
}

var possibleTiles = map[string]map[string][]string{
	"S": {
		"up":    []string{"|", "7", "F"},
		"right": []string{"-", "J", "7"},
		"down":  []string{"|", "L", "J"},
		"left":  []string{"-", "F", "L"},
	},
	"|": {
		"up":    []string{"|", "F", "7", "S"},
		"right": nil,
		"down":  []string{"|", "L", "J", "S"},
		"left":  nil,
	},
	"L": {
		"up":    []string{"|", "F", "7", "S"},
		"right": []string{"-", "J", "7", "S"},
		"down":  nil,
		"left":  nil,
	},
	"J": {
		"up":    []string{"|", "F", "7", "S"},
		"right": nil,
		"down":  nil,
		"left":  []string{"-", "L", "F", "S"},
	},
	"7": {
		"up":    nil,
		"right": nil,
		"down":  []string{"|", "L", "J", "S"},
		"left":  []string{"-", "L", "F", "S"},
	},
	"F": {
		"up":    nil,
		"right": []string{"-", "J", "7", "S"},
		"down":  []string{"|", "L", "J", "S"},
		"left":  nil,
	},
	"-": {
		"up":    nil,
		"right": []string{"-", "J", "7", "S"},
		"down":  nil,
		"left":  []string{"-", "L", "F", "S"},
	},
}

func New(gridSlice [][]string) grid {
	g := grid{
		G:      gridSlice,
		XBound: len(gridSlice[0]) - 1,
		YBound: len(gridSlice) - 1,
	}
	g.GoToStart()

	return g
}

type grid struct {
	G      [][]string
	Pos    [2]int
	Prev   [2]int
	XBound int
	YBound int
}

func (g *grid) GoToStart() {
	for y, row := range g.G {
		for x, cell := range row {
			if cell == "S" {
				g.Pos = [2]int{x, y}
				g.Prev = [2]int{x, y}
				return
			}
		}
	}

	panic(errors.New("can't find start"))
}

func (g *grid) FindNext() [2]int {
	currX := g.Pos[0]
	currY := g.Pos[1]
	currChar := g.CurrChar()

	var nextX int
	var nextY int
	for dir, val := range possibleDirections {
		nextX = currX + val[0]
		nextY = currY + val[1]
		if nextX < 0 || nextX > g.XBound {
			continue
		}

		if nextY < 0 || nextY > g.YBound {
			continue
		}

		if nextX == g.Prev[0] && nextY == g.Prev[1] {
			continue
		}

		nextChar := g.G[nextY][nextX]
		if nextChar == "." {
			continue
		}

		if slices.Contains(possibleTiles[currChar][dir], nextChar) {
			break
		}
	}

	return [2]int{nextX, nextY}
}

func (g *grid) CurrChar() string {
	return g.CharAt(g.Pos[0], g.Pos[1])
}

func (g *grid) CharAt(x, y int) string {
	// i just dont want to deal with bounds checking!
	if x < 0 || x > g.XBound {
		return "."
	}
	if y < 0 || y > g.YBound {
		return "."
	}

	return g.G[y][x]
}

func (g *grid) MoveToNext() {
	next := g.FindNext()
	g.Prev = g.Pos
	g.Pos = [2]int{next[0], next[1]}
}

func (g *grid) IsSVertex() bool {
	// technically i should save where we're at and then restore that before returning
	// but i'm not going to
	g.GoToStart()

	sPos := g.Pos
	first := g.FindNext()
	g.Prev = first
	second := g.FindNext()

	firstDiff := [2]int{first[0] - sPos[0], first[1] - sPos[1]}
	secondDiff := [2]int{second[0] - sPos[0], second[1] - sPos[1]}
	firstName := ""
	secondName := ""
	for name, pos := range possibleDirections {
		if pos == firstDiff {
			firstName = name
		}
		if pos == secondDiff {
			secondName = name
		}
	}
	if firstName == "" || secondName == "" {
		panic(errors.New("couldn't find nodes around S"))
	}

	if (firstName == "up" || firstName == "down") && (secondName == "up" || secondName == "down") {
		return false
	} else {
		return true
	}
}

func Shoelace(verticies [][2]int) int {
	result := 0
	lenV := len(verticies)
	for i := 1; i < lenV; i++ {
		x1 := verticies[i-1][0]
		y1 := verticies[i-1][1]
		x2 := verticies[i][0]
		y2 := verticies[i][1]

		p := (x1 * y2) - (x2 * y1)
		result += p
	}

	x1 := verticies[lenV-1][0]
	y1 := verticies[lenV-1][1]
	x2 := verticies[0][0]
	y2 := verticies[0][1]
	p := (x1 * y2) - (x2 * y1)
	result += p

	return result / 2
}

func PicksTheorem(area, points int) int {
	return area - (points / 2) + 1
}

func IsVertex(char string) bool {
	verticies := []string{"L", "J", "7", "F"}
	return slices.Contains(verticies, char)
}
