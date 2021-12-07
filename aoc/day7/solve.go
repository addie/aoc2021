package day7

import (
	"aoc2021/aoc"
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	aoc.Solver
}

func (s *Solver) Solve() int {
	data, err := parseFile(s.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	var res int
	switch s.Part {
	case 1:
		res = part1(data)
	case 2:
		res = part2(data)
	}

	return res
}

func part1(input []int) int {
	max := getMax(input)
	minCost := math.MaxInt
	for i := 0; i <= max; i++ {
		cost := calcCostSimple(input, i)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func part2(input []int) int {
	max := getMax(input)
	minCost := math.MaxInt
	for i := 0; i <= max; i++ {
		cost := calcCostComplex(input, i)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func calcCostSimple(input []int, idx int) int {
	totalCost := 0
	for _, in := range input {
		cost := math.Abs(float64(in - idx))
		totalCost += int(cost)
	}
	return totalCost
}

func calcCostComplex(input []int, idx int) int {
	totalCost := 0
	for _, in := range input {
		dist := int(math.Abs(float64(in - idx)))
		cost := dist * (dist + 1) / 2
		totalCost += int(cost)
	}
	return totalCost
}

func getMax(input []int) int {
	max := 0
	for _, in := range input {
		if in > max {
			max = in
		}
	}
	return max
}

func parseFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		parsedRow := strings.Split(row, ",")
		for _, n := range parsedRow {
			num, _ := strconv.Atoi(n)
			input = append(input, num)
		}
	}
	return input, nil
}
