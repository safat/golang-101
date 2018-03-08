package main

import (
	"fmt"
	"math"
)

type point struct {
	row, col int
}

var offsets = []point{
	{0, 0},
	{0, 1},
	{0, 2},
	{1, 1},
	{2, 0},
	{2, 1},
	{2, 2},
}

func main() {
	var n = 6
	var maxHourGlassValue = math.MinInt32

	grid := make([][]int, n)

	for i := range grid {
		grid[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			hourGlassValue := 0

			for _, point := range offsets {
				hourGlassValue += grid[i+point.row][j+point.col]
			}

			maxHourGlassValue = int(math.Max(float64(maxHourGlassValue), float64(hourGlassValue)))
		}
	}

	fmt.Println(maxHourGlassValue)
}

//https://www.hackerrank.com/challenges/2d-array/problem
