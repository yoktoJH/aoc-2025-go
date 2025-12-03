package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findMax(line string) (int, int) {
	maximum, imaximum := int32(0), 0
	for i, c := range line {
		intC := c - 48
		if intC > maximum {
			maximum = intC
			imaximum = i
		}
	}
	return int(maximum), imaximum
}

func part1(line string) int {
	maximum, imaximum := int32(0), 0
	for i, c := range line[:len(line)-1] {
		intC := c - 48
		if intC > maximum {
			maximum = intC
			imaximum = i
		}
	}
	submaximum := int32(0)
	for _, c := range line[imaximum+1:] {
		intC := c - 48
		if intC > submaximum {
			submaximum = intC

		}
	}
	fmt.Println(int(maximum)*10 + int(submaximum))
	return int(maximum)*10 + int(submaximum)
}

func part2(line string) int {
	sum := 0
	for i := 0; i < 12; i++ {
		m, ind := findMax(line[:len(line)-(11-i)])
		sum = sum*10 + m
		line = line[ind+1:]
	}
	fmt.Println(sum)
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
		line := scanner.Text()
		sum1 += part1(line)
		sum2 += part2(line)
	}
	fmt.Println("part1")
	fmt.Println(sum1)
	fmt.Println("part2")
	fmt.Println(sum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
