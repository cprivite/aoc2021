package main

import (
	"fmt"
	"log"
	"math"
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

	input, err := i.Strings(2021, 7)
	if err != nil {
		log.Fatal(err)
	}

	crabs := []int{}

	for _, line := range input {
		crabStrings := strings.Split(line, ",")
		for _, crabString := range crabStrings {
			crabInt, _ := strconv.Atoi(string(crabString))
			crabs = append(crabs, crabInt)
		}
	}
	sum := make([]float64, maxintSlice(crabs)+1)
	for i := 0; i <= maxintSlice(crabs); i++ {
		for _, j := range crabs {
			sum[i] += math.Abs(float64(i) - float64(j))
		}

	}
	fmt.Println("Part One")
	fuel, xpos := minintSlice(sum)
	fmt.Println("Fuel Cost:", fuel, "Position", xpos)

	sum2 := make([]float64, maxintSlice(crabs)+1)
	for destPos := 0; destPos <= maxintSlice(crabs); destPos++ {
		for _, crabPos := range crabs {
			x := math.Abs(float64(destPos) - float64(crabPos))
			sum2[destPos] += (x * (x + 1)) / 2
		}
	}
	fmt.Println("Part One")
	fuel, xpos = minintSlice(sum2)
	fmt.Println("Fuel Cost:", int(fuel), "Position", xpos)

}

func maxintSlice(slice []int) int {

	m := 0
	for i, e := range slice {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func minintSlice(slice []float64) (float64, int) {
	m := float64(0)
	mindex := 0
	for i, e := range slice {
		if i == 0 || e < m {
			m = e
			mindex = i
		}
	}
	return m, mindex
}
