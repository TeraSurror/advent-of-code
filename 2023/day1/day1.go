package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	numText := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file!")
		os.Exit(1)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		start := 0

		var num1 int
		var num2 int

		for start <= len(line)-1 {
			if len(line)-start >= 3 {
				value3, exists3 := numText[string(line[start:start+3])]
				if exists3 {
					num1 = value3
					break
				}
			}
			if len(line)-start >= 4 {
				value4, exists4 := numText[string(line[start:start+4])]
				if exists4 {
					num1 = value4
					break
				}
			}
			if len(line)-start >= 5 {
				value5, exists5 := numText[string(line[start:start+5])]
				if exists5 {
					num1 = value5
					break
				}
			}
			if line[start] >= '0' && line[start] <= '9' {
				num, _ := strconv.Atoi(string(line[start]))
				num1 = num
				break
			} else {
				start += 1
			}
		}
		start = 0
		for start <= len(line)-1 {
			if len(line)-start >= 3 {
				value3, exists3 := numText[string(line[start:start+3])]
				if exists3 {
					num2 = value3
				}
			}
			if len(line)-start >= 4 {
				value4, exists4 := numText[string(line[start:start+4])]
				if exists4 {
					num2 = value4
				}
			}
			if len(line)-start >= 5 {
				value5, exists5 := numText[string(line[start:start+5])]
				if exists5 {
					num2 = value5
				}
			}
			if line[start] >= '0' && line[start] <= '9' {
				num, _ := strconv.Atoi(string(line[start]))
				num2 = num
			}
			start += 1

		}
		fmt.Println(num1, num2)
		currNumber := (num1 * 10) + num2

		total += currNumber
	}

	fmt.Println(total)

}
