package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input-user.txt")
	// input := readFile("input-example.txt")
	input = input[:len(input)-1]

	fmt.Println("Part 1:", part1F(input))
	fmt.Println("Part 2:", part2F(input))
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}

func part1F(input []string) int {
	result := 0
	for _, line := range input {
		numStrs := strings.Split(line, " ")
		nums := make([][]int, 1)
		nums[0] = make([]int, len(numStrs))
		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums[0][i] = num
		}

		currRow := nums[0]                    // represents the row being iterated through
		newRow := make([]int, len(currRow)-1) // where we'll keep the diffs
		allZero := true                       // whether all the diffs are 0 for the current row
		i := 1                                // cell in current row
		y := 0                                // row number we're on
		for {
			prev := currRow[i-1]
			curr := currRow[i]

			diff := curr - prev
			allZero = allZero && diff == 0

			newRow[i-1] = diff
			i++
			if i == len(currRow) {
				nums = append(nums, newRow)
				if allZero {
					break
				}

				// reset for next row
				allZero = true
				i = 1
				y++
				currRow = nums[y]
				newRow = make([]int, len(currRow)-1)
			}
		}

		for i := len(nums) - 1; i > 0; i-- {
			row := nums[i]
			nextRow := nums[i-1]

			lastVal := row[len(row)-1]
			lastValNext := nextRow[len(nextRow)-1]
			newVal := lastVal + lastValNext

			nums[i-1] = append(nextRow, newVal)

			if i == 1 {
				result += newVal
			}
		}
	}

	return result
}

func part2F(input []string) int {
	result := 0
	for _, line := range input {
		numStrs := strings.Split(line, " ")
		nums := make([][]int, 1)
		nums[0] = make([]int, len(numStrs))
		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums[0][i] = num
		}

		currRow := nums[0]                    // represents the row being iterated through
		newRow := make([]int, len(currRow)-1) // where we'll keep the diffs
		allZero := true                       // whether all the diffs are 0 for the current row
		i := 1                                // cell in current row
		y := 0                                // row number we're on
		for {
			prev := currRow[i-1]
			curr := currRow[i]

			diff := curr - prev
			allZero = allZero && diff == 0

			newRow[i-1] = diff
			i++
			if i == len(currRow) {
				nums = append(nums, newRow)
				if allZero {
					break
				}

				// reset for next row
				allZero = true
				i = 1
				y++
				currRow = nums[y]
				newRow = make([]int, len(currRow)-1)
			}
		}

		for i := len(nums) - 1; i > 0; i-- {
			row := nums[i]
			nextRow := nums[i-1]

			firstVal := row[0]
			firstValNext := nextRow[0]
			newVal := firstValNext - firstVal

			nums[i-1] = append([]int{newVal}, nextRow...)

			if i == 1 {
				result += newVal
			}
		}
	}

	return result
}
