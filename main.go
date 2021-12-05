package main

import (
	"aoc2021/aoc"
	"aoc2021/aoc/day1"
	"aoc2021/aoc/day2"
	"aoc2021/aoc/day3"
	"aoc2021/aoc/day4"
	"aoc2021/aoc/day5"
	"aoc2021/data"
	"fmt"
	"log"
)

type version struct {
	day  int
	part int
}

func main() {

	// ============================================================
	v := version{day: 1, part: 1}
	day1Solver := &day1.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res := day1Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	v.part = 2
	day1Solver = &day1.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day1Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	// ============================================================
	v = version{day: 2, part: 1}
	day2Solver := day2.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day2Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	v.part = 2
	day2Solver = day2.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day2Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	// ============================================================
	v = version{day: 3, part: 1}
	day3Solver := day3.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day3Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	v.part = 2
	day3Solver = day3.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day3Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	// ============================================================
	v = version{day: 4, part: 1}
	day4Solver := day4.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day4Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	v.part = 2
	day4Solver = day4.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day4Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	// ============================================================
	// v = version{day: 5, part: 1}
	// day5Solver := day5.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	// res = day5Solver.Solve()
	// log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)

	v.day = 5
	v.part = 2
	day5Solver := day5.Solver{Solver: aoc.Solver{Part: v.part, FilePath: getData(v.day)}}
	res = day5Solver.Solve()
	log.Printf("%d.%d Answer: %d\n", v.day, v.part, res)
	// post(v.day, v.part, res)
}

func getData(day int) string {
	filename := fmt.Sprintf(data.Filename, day, data.Year)
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
