package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Set map[string]struct{}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func (s Set) Remove(item string) {
	delete(s, item)
}

func (s Set) Contains(item string) bool {
	_, exists := s[item]
	return exists
}

func main() {
	filePath := "input.txt"

	file, _ := os.Open(filePath)
	defer file.Close()

	// total := 0
	var pointsList []int
	var matchList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		currPoints, currMatches := getPoints(line)
		pointsList = append(pointsList, currPoints)
		matchList = append(matchList, currMatches)
	}

	var queue []int
	for i := 0; i < len(pointsList); i += 1 {
		queue = append(queue, i)
	}

	count := 0
	for len(queue) != 0 {
		cardNum := queue[0]
		count += 1
		// total += pointsList[cardNum]
		numCards := cardNum + matchList[cardNum]
		if numCards >= len(pointsList)-1 {
			numCards = len(pointsList) - 1
		}
		for i := cardNum + 1; i <= numCards; i += 1 {
			queue = append(queue, i)
		}
		queue = queue[1:]
	}

	fmt.Println(count)
}

func getPoints(line string) (int, int) {
	split1 := strings.Split(line, ":")
	allNumbers := strings.TrimSpace(split1[1])

	winningNumbersString := strings.TrimSpace(strings.Split(allNumbers, "|")[0])
	winningNumbersList := strings.Split(winningNumbersString, " ")
	winningNumbersSet := make(Set)
	for _, num := range winningNumbersList {
		if num == "" {
			continue
		}
		winningNumbersSet.Add(num)
	}

	ownedNumbersString := strings.TrimSpace(strings.Split(allNumbers, "|")[1])
	ownedNumbersList := strings.Split(ownedNumbersString, " ")

	matches := 0

	for _, num := range ownedNumbersList {
		if num == "" {
			continue
		}
		if winningNumbersSet.Contains(num) {
			matches += 1
		}
	}
	if matches == 0 {
		return 0, 0
	}
	points := math.Pow(2, float64(matches)-1)
	return int(points), matches
}
