/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package aoc

import "fmt"

func printGrid(grid [][]int) {
	for r := range grid {
		fmt.Print("[")
		for c := range grid[r] {
			fmt.Print(grid[r][c])
			fmt.Print(" ")
		}
		fmt.Print("]")
		fmt.Println()
	}
	fmt.Println()
}

func printGridBool(grid [][]bool) {
	for r := range grid {
		fmt.Print("[")
		for c := range grid[r] {
			if !grid[r][c] {
				fmt.Print(".")
			} else {
				fmt.Print(grid[r][c])
			}
			fmt.Print(" ")
		}
		fmt.Print("]")
		fmt.Println()
	}
	fmt.Println()
}
