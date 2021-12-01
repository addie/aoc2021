package main

import (
	"aoc2021/aoc/day1"
	"log"
)

func main() {
	err := day1.Solve(false, false)
	if err != nil {
		log.Fatal(err)
	}
}
