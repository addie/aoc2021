package day1

import (
	"aoc2021/aoc"
	"bufio"
	"log"
	"os"
	"strconv"
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

func part1(data []int) int {
	numIncreases := 0
	for i := range data {
		if i == 0 {
			continue
		}
		if data[i] > data[i-1] {
			numIncreases++
		}
	}
	return numIncreases
}

func part2(data []int) int {
	numIncreases := 0
	var window []int
	runningSum := 0
	for i := range data {
		if len(window) < 3 {
			window = append(window, data[i])
			runningSum += data[i]
			continue
		}
		nextSum := runningSum - window[0] + data[i]
		window = append(window[1:], data[i])
		if nextSum > runningSum {
			numIncreases++
		}
		runningSum = nextSum
	}
	return numIncreases
}

func parseFile(filePath string) ([]int, error) {
	var res []int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, digit)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		log.Fatal("read 0 bytes from file")
	}
	return res, nil
}
