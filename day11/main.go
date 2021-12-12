package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 11)
	if err != nil {
		log.Fatal(err)
	}

	octopi := [10][10]int{}
	for i, line := range input {
		for j, octopus := range line {
			octopi[i][j], _ = strconv.Atoi(string(octopus))
		}
	}
	flashCount := 0
	for z := 0; z < 300; z++ {
		flashed := false
		incCount := 0
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if octopi[y][x] == 0 {
					incCount++
				}
				octopi[y][x]++
			}
		}
		if incCount == 100 {
			fmt.Println("ALL FLASHED", z)
		}

		flashed = processFlashed(&octopi, &flashCount)

		for flashed {
			flashed = processFlashed(&octopi, &flashCount)
		}

	}
	for i := range octopi {
		fmt.Println(octopi[i])
	}
	fmt.Println(flashCount)
}

func processFlashed(octopi *[10][10]int, flashCount *int) bool {
	flashed := false
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if octopi[y][x] > 9 {
				flashed = true
				*flashCount++
				octopi[y][x] = 0
				if y > 0 {
					if x > 0 {
						if octopi[y-1][x-1] != 0 {
							octopi[y-1][x-1]++
						}
					}
					if octopi[y-1][x] != 0 {
						octopi[y-1][x]++
					}
					if x < 9 {
						if octopi[y-1][x+1] != 0 {
							octopi[y-1][x+1]++
						}
					}
				}
				if y < 9 {
					if x > 0 {
						if octopi[y+1][x-1] != 0 {
							octopi[y+1][x-1]++
						}
					}
					if octopi[y+1][x] != 0 {
						octopi[y+1][x]++
					}
					if x < 9 {
						if octopi[y+1][x+1] != 0 {
							octopi[y+1][x+1]++
						}
					}
				}
				if x > 0 {
					if octopi[y][x-1] != 0 {
						octopi[y][x-1]++
					}
				}
				if x < 9 {
					if octopi[y][x+1] != 0 {
						octopi[y][x+1]++
					}
				}
			}
		}
	}
	return flashed
}
