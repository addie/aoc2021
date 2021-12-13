package main

import (
	"aoc2021/aoc"
	"aoc2021/clipboard"
	"aoc2021/data"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type version struct {
	day  int
	part int
}

func main() {
	problemSet := os.Args[1:][0]
	dayStr := strings.Split(problemSet, ".")[0]
	day, _ := strconv.Atoi(dayStr)
	levelStr := strings.Split(problemSet, ".")[1]
	level, _ := strconv.Atoi(levelStr)

	saveData(day)
	res := 0
	switch problemSet {
	case "1.1":
		res = aoc.Day1Part1()
	case "1.2":
		res = aoc.Day1Part2()
	case "2.1":
		res = aoc.Day2Part1()
	case "2.2":
		res = aoc.Day2Part2()
	case "3.1":
		res = aoc.Day3Part1()
	case "3.2":
		res = aoc.Day3Part2()
	case "4.1":
		res = aoc.Day4Part1()
	case "4.2":
		res = aoc.Day4Part2()
	case "5.1":
		res = aoc.Day5Part1()
	case "5.2":
		res = aoc.Day5Part2()
	case "6.1":
		res = aoc.Day6Part1()
	case "6.2":
		res = aoc.Day6Part2()
	case "7.1":
		res = aoc.Day7Part1()
	case "7.2":
		res = aoc.Day7Part2()
	case "8.1":
		res = aoc.Day8Part1()
	case "8.2":
		res = aoc.Day8Part2()
	case "9.1":
		res = aoc.Day9Part1()
	case "9.2":
		res = aoc.Day9Part2()
	case "10.1":
		res = aoc.Day10Part1()
	case "10.2":
		res = aoc.Day10Part2()
	case "11.1":
		res = aoc.Day11Part1()
	case "11.2":
		res = aoc.Day11Part2()
	case "12.1":
		res = aoc.Day12Part1()
	case "12.2":
		res = aoc.Day12Part2()
	}
	fmt.Printf("Problem %d.%d - Answer %d\n", day, level, res)
	clipboard.WriteAll(strconv.Itoa(res))
	// post(day, level, res)
}

func saveData(day int) string {
	filename := fmt.Sprintf(data.Filename, day)
	err := data.Get(day, filename)
	if err != nil {
		log.Fatal(err)
	}
	return filename
}

func post(day, level, res int) {
	err := data.Post(day, level, res)
	if err != nil {
		log.Fatal(err)
	}
}
