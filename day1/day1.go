package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	if err != nil {
		os.Exit(1)
	}

	pings := make([]int, 0)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		pings = append(pings, number)
	}

	err = file.Close()
	if err != nil {
		os.Exit(1)
	}

	timesIncremented := 0
	lastPing := 0
	for i, ping := range pings {
		if i == 0 {
			timesIncremented++
			lastPing = ping
		} else if i < len(pings) {
			if ping > lastPing {
				timesIncremented++
			}
			lastPing = ping
		}
	}

	fmt.Println("Total times incremented: ", timesIncremented)

	timesIncrementedAgain := 0

	for i := 0; i+3 < len(pings); i++ {
		ping := pings[i]
		if ping < pings[i+3] {
			timesIncrementedAgain = timesIncrementedAgain + 1
		}
	}

	fmt.Println("Total times incremented: ", timesIncrementedAgain)
}
