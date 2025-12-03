package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path := "input"
	//path := "example"
	file, err := os.Open(path + "/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	count2 := 0
	arrow := 50
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		shift, err := strconv.Atoi(line[1:])
		count2 += shift / 100
		shift %= 100
		if err != nil {
			log.Fatal(err)
		}
		switch line[0] {
		case 'L':
			arrow = arrow - shift
			if arrow < 0 && arrow != -shift {
				count2 += 1
			}

		case 'R':
			arrow = arrow + shift
			count2 += arrow / 101
			//don't count cases where arrow==100
		}

		arrow %= 100
		//remainder stays negative
		if arrow < 0 {
			arrow += 100
		}
		if arrow == 0 {
			count++
			count2++
		}
	}

	fmt.Println("part1")
	fmt.Println(count)
	fmt.Println("part2")
	fmt.Println(count2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
