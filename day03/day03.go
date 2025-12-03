package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findMax(line []byte) (int, int) {
	maximum, imaximum := byte(0), 0
	for i, c := range line {
		intC := c - 48
		if intC > maximum {
			maximum = intC
			imaximum = i
		}
	}
	return int(maximum), imaximum
}

func solve(line []byte, digits int) int {
	sum := 0
	for i := 0; i < digits; i++ {
		maximum, imaximum := findMax(line[:len(line)-(digits-1-i)])
		sum = 10*sum + maximum
		line = line[imaximum+1:]
	}
	return sum
}

func main() {
	path := "input"
	//path := "example"
	file, err := os.Open(path + "/03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum1 := 0
	sum2 := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Bytes()
		sum1 += solve(line, 2)
		sum2 += solve(line, 12)
	}
	fmt.Println("part1")
	fmt.Println(sum1)
	fmt.Println("part2")
	fmt.Println(sum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
