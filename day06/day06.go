package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readColum(numbers []string, column int) int {
	outOfBound := true
	for j := 0; j < len(numbers)-1; j++ {
		if column < len(numbers[j]) {
			outOfBound = false
		}
	}
	if outOfBound {
		return 0
	}
	num := 0
	for j := 0; j < len(numbers)-1; j++ {
		if column < len(numbers[j]) && numbers[j][column] != ' ' {
			num *= 10
			num += int(numbers[j][column] - '0')
		}
	}
	return num
}

func main() {
	path := "input"
	//path := "example"
	file, err := os.Open(path + "/06.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	input := make([][]string, 0)
	inputLines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		inputLines = append(inputLines, line)
		input = append(input, strings.Fields(line))
		fmt.Println(input)
	}
	results := make([]int, 0)

	for i := range input[0] {
		result := 0

		if input[len(input)-1][i] == "*" {
			result = 1
		}
		for j := 0; j < len(input)-1; j++ {
			x, _ := strconv.Atoi(input[j][i])
			if input[len(input)-1][i] == "*" {
				result *= x
			} else {
				result += x
			}
		}
		results = append(results, result)
	}
	sum := 0
	for _, x := range results {
		sum += x
	}

	results2 := make([]int, 0)
	i := 0
	for i < len(inputLines[0]) {
		op := inputLines[len(inputLines)-1][i]
		result := 0
		if op == '*' {
			result = 1
		}
		x := readColum(inputLines, i)
		i++
		fmt.Println(x)
		for x != 0 {
			fmt.Println(x)
			if op == '*' {
				result *= x
			} else {
				result += x
			}
			x = readColum(inputLines, i)
			i++

		}
		results2 = append(results2, result)
	}
	sum2 := 0
	for _, x := range results2 {
		sum2 += x
	}
	fmt.Println("part1")
	fmt.Println(sum)
	fmt.Println("part2")
	fmt.Println(sum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
