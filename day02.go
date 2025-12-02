package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(line string) int {
	ranges := strings.Split(line, ",")
	sum := 0

	for _, interval := range ranges {
		numbers := strings.Split(interval, "-")
		low := numbers[0]
		high := numbers[1]
		half := 0
		length := len(low) / 2
		if length == 0 {
			length = 1
		}
		//sum += try_len(i, low, high)
		tmp, err := strconv.Atoi(low[:length])
		half = tmp
		if err != nil {
			log.Fatal(err)
		}
		low_int, err := strconv.Atoi(low)
		if err != nil {
			log.Fatal(err)
		}
		high_int, err := strconv.Atoi(high)
		if err != nil {
			log.Fatal(err)
		}

		value := half + half*int(math.Pow(float64(10), math.Trunc(math.Log10(float64(half)))+1))
		for value <= high_int {
			if low_int <= value {
				fmt.Println(value)
				sum += value
			}
			half += 1
			value = half + half*int(math.Pow(float64(10), math.Trunc(math.Log10(float64(half)))+1))
		}
	}
	return sum
}

func main() {
	path := "input"
	//path := "example"
	file, err := os.Open(path + "/02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()
	line := scanner.Text()

	sum1 := solve(line)
	fmt.Println("part1")
	fmt.Println(sum1)
	fmt.Println("part2")
	fmt.Println()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
