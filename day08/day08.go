package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type box struct {
	x, y, z int
}

func findSmallestDistance(distances [][]float64) (int, int) {
	mind := math.Inf(1)
	mini := 0
	minj := 0
	for i := 0; i < len(distances); i++ {
		for j := 0; j < len(distances[0]); j++ {
			if distances[i][j] != -1 && distances[i][j] < mind {
				mind = distances[i][j]
				mini = i
				minj = j
			}
		}
	}
	//fmt.Println("connecting", mini, minj, mind)
	return mini, minj
}

func mergeComponents(arr1, arr2 []int) []int {
	tmp := make(map[int]int)
	for _, x := range arr1 {
		tmp[x] = 1
	}
	for _, x := range arr2 {
		tmp[x] = 1
	}
	result := make([]int, 0)
	for key := range tmp {
		result = append(result, key)
	}
	return result
}

func insertLength(topLen []int, val int) bool {
	inserted := false
	for i := 0; i < 3; i++ {
		if val > topLen[i] {
			tmp := topLen[i]
			topLen[i] = val
			val = tmp
			inserted = true
		}
	}
	return inserted
}

func makeNConnections(n int, distances [][]float64) int {
	connectedComponents := make(map[int][]int)
	for i := 0; i < len(distances); i++ {
		connectedComponents[i] = []int{i}
	}
	for i := 0; i < n; i++ {
		b1, b2 := findSmallestDistance(distances)
		distances[b1][b2] = -1
		distances[b2][b1] = -1
		newComponent := mergeComponents(connectedComponents[b1], connectedComponents[b2])
		for _, b := range newComponent {
			connectedComponents[b] = newComponent
		}
	}
	topLen := make([]int, 4)
	for _, comp := range connectedComponents {
		if insertLength(topLen, len(comp)) {
			for _, x := range comp {
				connectedComponents[x] = make([]int, 0)
			}
		}
	}
	return topLen[0] * topLen[1] * topLen[2]
}

func makeAllConnections(boxes []box, distances [][]float64) int {
	connectedComponents := make(map[int][]int)
	for i := 0; i < len(distances); i++ {
		connectedComponents[i] = []int{i}
	}
	lasCompLen := 1
	lastXMult := 0
	for lasCompLen < len(boxes) {
		b1, b2 := findSmallestDistance(distances)
		distances[b1][b2] = -1
		distances[b2][b1] = -1
		if slices.Contains(connectedComponents[b1], b2) {
			continue
		}

		newComponent := mergeComponents(connectedComponents[b1], connectedComponents[b2])
		lasCompLen = len(newComponent)
		for _, b := range newComponent {
			connectedComponents[b] = newComponent
		}
		lastXMult = boxes[b1].x * boxes[b2].x
	}
	return lastXMult
}

func main() {
	path := "input"
	//path := "example"
	file, err := os.Open(path + "/08.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	boxes := make([]box, 0)

	distances := make([][]float64, 0)
	distances2 := make([][]float64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		z, err := strconv.Atoi(nums[2])
		if err != nil {
			log.Fatal(err)
		}
		boxes = append(boxes, box{x, y, z})
	}
	for i := 0; i < len(boxes); i++ {
		oneBoxDist := make([]float64, 0)
		oneBoxDist2 := make([]float64, 0)
		for j := 0; j < len(boxes); j++ {
			if i == j {
				oneBoxDist = append(oneBoxDist, math.Inf(1))
				oneBoxDist2 = append(oneBoxDist2, math.Inf(1))
				continue
			}
			box1 := boxes[i]
			box2 := boxes[j]
			xdif := box1.x - box2.x
			ydif := box1.y - box2.y
			zdif := box1.z - box2.z
			oneBoxDist = append(oneBoxDist, math.Sqrt(float64(xdif*xdif+ydif*ydif+zdif*zdif)))
			oneBoxDist2 = append(oneBoxDist2, math.Sqrt(float64(xdif*xdif+ydif*ydif+zdif*zdif)))
		}
		distances = append(distances, oneBoxDist)
		distances2 = append(distances2, oneBoxDist2)
	}

	fmt.Println("part1")
	fmt.Println(makeNConnections(1000, distances))
	fmt.Println("part2")
	fmt.Println(makeAllConnections(boxes, distances2))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
