package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(input []string) (int, int) {
	beams := make(map[int]int)
	splits := 0
	ibeam := 0
	for i := 0; i < len(input[0]); i++ {
		if input[0][i] == 'S' {
			ibeam = i
			beams[i] = 1
		}
	}
	graph := make([]map[int]int, 0)
	graph = append(graph, beams)
	for row := 1; row < len(input); row++ {
		newBeams := make(map[int]int)
		for beam := range beams {
			if input[row][beam] == '^' {
				splits++
				newBeams[beam-1] = 1
				newBeams[beam+1] = 1
			} else {
				newBeams[beam] = 1
			}

		}
		graph = append(graph, newBeams)
		beams = newBeams
	}
	for i := len(input) - 2; i >= 0; i-- {
		for col, _ := range graph[i] {
			val, ok := graph[i+1][col]
			if ok {
				graph[i][col] = val
			} else {
				graph[i][col] = graph[i+1][col-1] + graph[i+1][col+1]
			}

		}
	}
	return splits, graph[0][ibeam]
}

func main() {
	path := "input"
	//path := "example"
	file, err := os.Open(path + "/07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	s1, s2 := part1(input)
	fmt.Println("part1")
	fmt.Println(s1)
	fmt.Println("part2")
	fmt.Println(s2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
