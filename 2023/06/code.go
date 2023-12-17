package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	Time     int
	Distance int
}

func main() {
	input := readFile("input-user.txt")
	races := parseInput(input)

	fmt.Println("Part 1:", part1f(races))
	fmt.Println("Part 2:", part2f(races))
}

func part1f(races []race) int {
	winning := make([]int, 0, len(races))
	for _, race := range races {
		winningOptions := 0
		for i := 1; i < race.Time; i++ {
			resultDistance := calculateDistance(race.Time, i)
			if resultDistance > race.Distance {
				winningOptions++
			}
		}
		winning = append(winning, winningOptions)
	}

	result := winning[0]
	for i := 1; i < len(winning); i++ {
		result *= winning[i]
	}

	return result
}

func part2f(races []race) int {
	timeStr := ""
	distanceStr := ""
	for _, race := range races {
		timeStr += strconv.Itoa(race.Time)
		distanceStr += strconv.Itoa(race.Distance)
	}
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatal(err)
	}

	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		log.Fatal(err)
	}

	lastResult := 0
	winningOptions := 0
	for i := 1; i < time; i++ {
		resultDistance := calculateDistance(time, i)
		if resultDistance > distance {
			winningOptions++
		}
		if resultDistance < distance && lastResult > resultDistance {
			break
		}
		lastResult = resultDistance
	}

	return winningOptions
}

func calculateDistance(time int, holdTime int) int {
	return holdTime * (time - holdTime)
}

func parseInput(input []string) []race {
	timeLine := strings.Split(input[0], "Time:")
	timesStr := strings.Split(timeLine[1], " ")
	times := make([]int, 0)
	for _, timeStr := range timesStr {
		if timeStr == "" {
			continue
		}
		i, err := strconv.Atoi(strings.TrimSpace(timeStr))
		if err != nil {
			log.Fatal(err)
		}
		times = append(times, i)
	}

	distanceLine := strings.Split(input[1], "Distance:")
	distancesStr := strings.Split(distanceLine[1], " ")
	distances := make([]int, 0)
	for _, distanceStr := range distancesStr {
		if distanceStr == "" {
			continue
		}
		i, err := strconv.Atoi(strings.TrimSpace(distanceStr))
		if err != nil {
			log.Fatal(err)
		}
		distances = append(distances, i)
	}

	races := make([]race, 0, len(times))
	for i, time := range times {
		r := race{
			Time:     time,
			Distance: distances[i],
		}
		races = append(races, r)
	}

	return races
}

func readFile(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\r\n")
}
