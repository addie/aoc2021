package day6

import (
	"aoc2021/aoc"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Solver struct {
	aoc.Solver
}

func (s *Solver) Solve() int {
	coords, err := parseFile(s.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	var res int
	switch s.Part {
	case 1:
		res = part1(coords)
	case 2:
		res = part2(coords)
	}

	return res
}

func part1(input []int) int {
	const totalDays = 80
	return lanternFishCount(totalDays, input)
}

func part2(input []int) int {
	const totalDays = 256
	return lanternFishCount(totalDays, input)
}

func lanternFishCount(days int, input []int) int {
	const resetFish = 6
	const newFish = 8
	counter := counter(input)
	for day := 0; day < days; day++ {
		fishMap := make(map[int]int) // tracks how many fish have n numbers to give birth
		for numDays, count := range counter {
			if numDays == 0 {
				fishMap[resetFish] += count
				fishMap[newFish] += count
			} else {
				fishMap[numDays-1] += count
			}
		}
		counter = fishMap
	}
	return sum(counter)
}

func sum(counter map[int]int) int {
	count := 0
	for _, val := range counter {
		count += val
	}
	return count
}

func counter(input []int) map[int]int {
	c := make(map[int]int)
	for _, in := range input {
		if _, ok := c[in]; ok {
			c[in]++
		} else {
			c[in] = 1
		}
	}
	return c
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
