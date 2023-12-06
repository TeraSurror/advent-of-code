package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		powerSet := isGamePossible(line)
		total += powerSet
	}

	fmt.Println(total)
}

func isGamePossible(line string) int {
	split1 := strings.Split(line, ":")

	split2 := strings.Split(split1[1], ";")

	minRed := 0
	minBlue := 0
	minGreen := 0

	for _, game := range split2 {
		game = strings.TrimSpace(game)
		cubes := strings.Split(game, ",")

		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			cube_dets := strings.Split(cube, " ")
			cube_count, _ := strconv.Atoi(cube_dets[0])
			if cube_dets[1] == "red" {
				if cube_count > minRed {
					minRed = cube_count
				}
			} else if cube_dets[1] == "blue" {
				if cube_count > minBlue {
					minBlue = cube_count
				}
			} else if cube_dets[1] == "green" {
				if cube_count > minGreen {
					minGreen = cube_count
				}
			}
		}
	}
	return minBlue * minGreen * minRed
}
