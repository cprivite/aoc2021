package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/echojc/aocutil"
)

type line struct {
	startx int
	starty int
	endx   int
	endy   int
}

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 5)
	if err != nil {
		log.Fatal(err)
	}

	var findInt = regexp.MustCompile(`\d+`)
	lines := make([]line, len(input))
	for i, s := range input {
		t := findInt.FindAllString(s, -1)
		lines[i].startx, _ = strconv.Atoi(string(t[0]))
		lines[i].starty, _ = strconv.Atoi(string(t[1]))
		lines[i].endx, _ = strconv.Atoi(string(t[2]))
		lines[i].endy, _ = strconv.Atoi(string(t[3]))
	}

	myMap := make([][]int, 1000)
	for i := range myMap {
		myMap[i] = make([]int, 1000)
	}

	for _, line := range lines {
		myMap = drawLine(myMap, line.startx, line.starty, line.endx, line.endy)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	for _, line := range myMap {
		fmt.Println(line)
	}
	fmt.Println(countPointsOverTwo(myMap))
}

func countPointsOverTwo(myMap [][]int) int {
	count := 0
	for _, i := range myMap {
		for _, j := range i {
			if j > 1 {
				count++
			}
		}
	}
	return count
}

func drawLine(myMap [][]int, startx int, starty int, endx int, endy int) [][]int {
	if startx == endx {
		for i := min(starty, endy); i <= max(starty, endy); i++ {
			myMap[i][startx] += 1
		}
	} else if starty == endy {
		for i := min(startx, endx); i <= max(startx, endx); i++ {
			myMap[starty][i] += 1
		}
	} else if startx < endx && starty < endy {
		for i, j := startx, starty; i <= endx; i++ {
			myMap[j][i] += 1
			j++
		}
	} else if startx < endx && starty > endy {
		for i, j := startx, starty; i <= endx; i++ {
			myMap[j][i] += 1
			j--
		}
	} else if startx > endx && starty < endy {
		for i, j := startx, starty; j <= endy; j++ {
			myMap[j][i] += 1
			i--
		}
	} else if startx > endx && starty > endy {
		for i, j := startx, starty; i >= endx; i-- {
			myMap[j][i] += 1
			j--
		}
	}
	return myMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
