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
		res = aoc.P11()
	case "1.2":
		res = aoc.P12()
	case "2.1":
		res = aoc.P21()
	case "2.2":
		res = aoc.P22()
	case "3.1":
		res = aoc.P31()
	case "3.2":
		res = aoc.P32()
	case "4.1":
		res = aoc.P41()
	case "4.2":
		res = aoc.P42()
	case "5.1":
		res = aoc.P51()
	case "5.2":
		res = aoc.P52()
	case "6.1":
		res = aoc.P61()
	case "6.2":
		res = aoc.P62()
	case "7.1":
		res = aoc.P71()
	case "7.2":
		res = aoc.P72()
	case "8.1":
		res = aoc.P81()
	case "8.2":
		res = aoc.P82()
	case "9.1":
		res = aoc.P91()
	case "9.2":
		res = aoc.P92()
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
