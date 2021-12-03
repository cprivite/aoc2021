package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "input.txt"

	fileBytes, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pings := strings.Split(string(fileBytes), "\n")

	var timesIncremented int
	var lastPing string

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
}
