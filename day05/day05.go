package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func combineRanges(ranges [][]int) ([][]int, bool) {
	newRanges := make([][]int, 0)
	changed := false
	for _, r := range ranges {
		for _, newR := range newRanges {
			if r[0] > r[1] {
				break
			}
			if newR[0] <= r[0] && r[1] <= newR[1] {
				changed = true
				r[0] = 1
				r[1] = 0
				break
			} else if r[0] <= newR[0] && newR[1] <= r[1] {
				changed = true
				newR[0] = r[0]
				newR[1] = r[1]
				r[0] = 1
				r[1] = 0
				break
			} else if r[0] <= newR[0] && newR[0] <= r[1] && r[1] <= newR[1] {
				r[1] = newR[0] - 1
				changed = true
			} else if newR[0] <= r[0] && r[0] <= newR[1] && newR[1] <= r[1] {
				r[0] = newR[1] + 1
				changed = true
			}
		}
		if r[0] <= r[1] {
			newRanges = append(newRanges, []int{r[0], r[1]})
		}

	}
	return newRanges, changed
}

func main() {
	path := "input"
	//path := "example"
	file, err := os.Open(path + "/05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := make([][]int, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		numbers := strings.Split(line, "-")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		ranges = append(ranges, []int{x, y})

	}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		x, _ := strconv.Atoi(line)

		for _, r := range ranges {
			if x <= r[1] && x >= r[0] {
				//fmt.Println(x)
				sum++
				break
			}
		}
	}
	changed := true
	for changed {
		changed = false
		ranges, changed = combineRanges(ranges)
	}
	sum2 := 0
	for _, r := range ranges {
		sum2 += r[1] - r[0] + 1
	}

	fmt.Println("part1")
	fmt.Println(sum)
	fmt.Println("part2")
	fmt.Println(sum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
