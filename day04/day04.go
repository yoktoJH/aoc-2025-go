package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countNeighbours(i, j int, rolls [][]byte) int {
	sum := 0
	for iShift := -1; iShift < 2; iShift++ {
		ii := i + iShift
		if ii < 0 || ii >= len(rolls) {
			continue
		}
		for jShift := -1; jShift < 2; jShift++ {
			if iShift == 0 && jShift == 0 {
				continue
			}
			jj := j + jShift
			if jj < 0 || jj >= len(rolls[ii]) {
				continue
			}
			if rolls[ii][jj] == 64 {
				sum++
			}
		}
	}
	return sum
}

func removeOnRepeat(rolls [][]byte, toRemove [][]int) int {
	var sum int
	changed := true
	for changed {
		sum += len(toRemove)
		changed = false
		for _, p := range toRemove {
			rolls[p[0]][p[1]] = 46
		}
		toRemove = toRemove[:0]
		for i := 0; i < len(rolls); i++ {
			for j := 0; j < len(rolls[i]); j++ {

				if rolls[i][j] == 64 && countNeighbours(i, j, rolls) < 4 {
					toRemove = append(toRemove, []int{i, j})
				}
			}
		}
		changed = len(toRemove) != 0
	}
	return sum
}

func main() {
	//path := "input"
	path := "example"
	file, err := os.Open(path + "/04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	rolls := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		rolls = append(rolls, ([]byte)(line))
	}
	sum := 0
	toRemove := make([][]int, 0)
	for i := 0; i < len(rolls); i++ {
		for j := 0; j < len(rolls[i]); j++ {
			if rolls[i][j] == 64 && countNeighbours(i, j, rolls) < 4 {
				sum++
				toRemove = append(toRemove, []int{i, j})
			}
		}
	}

	fmt.Println("part1")
	fmt.Println(sum)
	fmt.Println("part2")
	fmt.Println(removeOnRepeat(rolls, toRemove))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
