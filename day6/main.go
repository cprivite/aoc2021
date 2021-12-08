package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 6)
	if err != nil {
		log.Fatal(err)
	}

	fishes := []int{}

	for _, line := range input {
		fishStrings := strings.Split(line, ",")
		for _, fishString := range fishStrings {
			fishInt, _ := strconv.Atoi(string(fishString))
			fishes = append(fishes, fishInt)
		}
	}

	twoFishes := make([]int, len(fishes))
	copy(twoFishes, fishes)

	/* 	for j := 0; j < 80; j++ {
		fishCount := len(fishes)
		for i := 0; i < fishCount; i++ {
			if fishes[i] == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i] = fishes[i] - 1
			}
		}
	} */

	fishDex := [9]int{}

	for _, twoFish := range twoFishes {
		fishDex[twoFish] += 1
	}

	for j := 0; j < 256; j++ {
		fishSpawn := fishDex[0]
		for i := 0; i < 8; i++ {
			fishDex[i] = fishDex[i+1]
		}
		fishDex[8] = fishSpawn
		fishDex[6] += fishSpawn
	}

	fishCount := 0
	for _, i := range fishDex {
		fishCount += i
	}

	fmt.Println("fishCount:", fishCount)

}
