package day2

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

type pos struct {
	horizontal, depth, aim int
}

func part1(data [][]string) int {
	res := pos{}
	for _, fullCmd := range data {
		dir := fullCmd[0]
		num, _ := strconv.Atoi(fullCmd[1])
		switch dir {
		case "forward":
			res.horizontal += num
		case "down":
			res.depth += num
		case "up":
			res.depth -= num
		}
	}
	return res.horizontal * res.depth
}

func part2(data [][]string) int {
	res := pos{}
	for _, fullCmd := range data {
		dir := fullCmd[0]
		num, _ := strconv.Atoi(fullCmd[1])
		switch dir {
		case "forward":
			res.horizontal += num
			res.depth += res.aim * num
		case "down":
			res.aim += num
		case "up":
			res.aim -= num
		}
	}
	return res.horizontal * res.depth
}

func parseFile(filePath string) ([][]string, error) {
	var res [][]string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fullCmd := strings.Split(scanner.Text(), " ")
		dir, numStr := fullCmd[0], fullCmd[1]
		res = append(res, []string{dir, numStr})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		log.Fatal("read 0 bytes from file")
	}
	return res, nil
}
