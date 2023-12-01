package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parseNumbersPart1(line string) (int, int) {
	parsedNumbers := []int{}
	for _, c := range line {
		num, err := strconv.Atoi(string(c))
		if err == nil {
			parsedNumbers = append(parsedNumbers, num)
		}
	}

	return parsedNumbers[0], parsedNumbers[len(parsedNumbers)-1]
}

func parseNumbersPart2(line string) (int, int) {
	parsedNumbers := []int{}
	for i, c := range line {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			for w, d := range numberMap {
				if strings.HasPrefix(line[i:], w) {
					parsedNumbers = append(parsedNumbers, d)
				}
			}
		} else {
			parsedNumbers = append(parsedNumbers, num)
		}
	}

	return parsedNumbers[0], parsedNumbers[len(parsedNumbers)-1]
}

func main() {
	myfile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could open file: %s", err)
	}
	defer myfile.Close()

	sum_1 := 0
	sum_2 := 0
	reader := bufio.NewScanner(myfile)

	for reader.Scan() {
		line := reader.Text()

		first_num1, last_num1 := parseNumbersPart1(line)
		sum_1 += first_num1*10 + last_num1

		first_num2, last_num2 := parseNumbersPart2(line)
		sum_2 += first_num2*10 + last_num2
	}

	fmt.Println("Day 1")
	fmt.Printf("part1: %d\n", sum_1)
	fmt.Printf("part2: %d\n", sum_2)
}
