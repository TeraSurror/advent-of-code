package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := "input.txt"

	file, _ := os.Open(filePath)
	defer file.Close()

	var inputGrid [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var inputGridRow []string
		for _, letter := range line {
			inputGridRow = append(inputGridRow, string(letter))
		}

		inputGrid = append(inputGrid, inputGridRow)

	}

	result := findSumOfPartNums(inputGrid)

	fmt.Println(result)
}

func findSumOfPartNums(grid [][]string) int {
	star_map := make(map[string][]int)
	total := 0
	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if grid[i][j] >= "0" && grid[i][j] <= "9" {
				num := ""
				k := j
				for j < len(grid[0]) && grid[i][j] >= "0" && grid[i][j] <= "9" {
					num += grid[i][j]
					j += 1
				}
				findIfNumNextToSymbol(grid, i, k, j, star_map, num)
				// if isNextToSymbol {
				// 	num1, _ := strconv.Atoi(num)
				// 	total += num1
				// }
			}
		}
	}

	for _, valueList := range star_map {
		if len(valueList) > 1 {
			currValue := 1
			for _, val := range valueList {
				currValue *= val
			}
			total += currValue
		}
	}

	return total
}

func findIfNumNextToSymbol(grid [][]string, i int, start int, end int, star_map map[string][]int, num string) bool {
	directions := [][]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	for j := start; j < end; j += 1 {
		for _, direction := range directions {
			row := i + direction[0]
			col := j + direction[1]

			if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) {
				if !(grid[row][col] == "." || (grid[row][col] >= "0" && grid[row][col] <= "9")) {
					if grid[row][col] == "*" {
						key := strconv.Itoa(row) + "-" + strconv.Itoa(col)
						num1, _ := strconv.Atoi(num)
						star_map[key] = append(star_map[key], num1)
					}
					return true
				}
			}
		}
	}

	return false
}
