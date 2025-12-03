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

func computeVal(subNum, length, low int) int {
	shift := int(math.Pow(float64(10), float64(length)))
	value := subNum*shift + subNum
	for value < low {
		value = value*shift + subNum
	}
	return value
}

func part2(numString string, low, high int) int {
	sum := 0
	knownValues := make(map[int]int)
	for length := 1; length <= len(numString)/2; length++ {
		baseVal := int(math.Pow(float64(10), float64(length-1)))
		maxVal := int(math.Pow(float64(10), float64(length)))
		for i := baseVal; i < maxVal; i++ {
			value := computeVal(i, length, low)
			if _, contains := knownValues[value]; !contains && value <= high {
				fmt.Println(value)
				sum += value
				knownValues[value] = 1
			}
		}
	}

	return sum
}

func part1(half, low, high int) int {
	sum := 0
	value := half + half*int(math.Pow(float64(10), math.Trunc(math.Log10(float64(half)))+1))
	for value <= high {
		if low <= value {
			sum += value
		}
		half += 1
		value = half + half*int(math.Pow(float64(10), math.Trunc(math.Log10(float64(half)))+1))
	}
	return sum
}

func solve(line string) (int, int) {
	ranges := strings.Split(line, ",")
	sum1 := 0
	sum2 := 0
	for _, interval := range ranges {
		fmt.Println(interval)
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
		lowInt, err := strconv.Atoi(low)
		if err != nil {
			log.Fatal(err)
		}
		highInt, err := strconv.Atoi(high)
		if err != nil {
			log.Fatal(err)
		}
		sum1 += part1(half, lowInt, highInt)
		sum2 += part2(high, lowInt, highInt)

	}
	return sum1, sum2
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

	sum1, sum2 := solve(line)
	fmt.Println("part1")
	fmt.Println(sum1)
	fmt.Println("part2")
	fmt.Println(sum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
