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

	input, err := i.Strings(2021, 13)
	if err != nil {
		log.Fatal(err)
	}

	dots := make([][2]int, 0)

	for _, line := range input {
		s := [2]int{}
		s[0], _ = strconv.Atoi(strings.Split(line, ",")[0])
		s[1], _ = strconv.Atoi(strings.Split(line, ",")[1])
		dots = append(dots, s)
	}
	maxx := 0
	maxy := 0
	for _, s := range dots {
		if s[1] > maxy {
			maxy = s[1]
		}

		if s[0] > maxx {
			maxx = s[0]
		}
		fmt.Println(s)
	}
	fmt.Println(maxx, maxy)

	paper := make([][]string, maxy+1)
	for y := range paper {
		paper[y] = make([]string, maxx+1)
	}

	for y := range paper {
		for x := range paper[y] {
			paper[y][x] = "."
		}
	}

	for _, d := range dots {
		paper[d[1]][d[0]] = "#"
	}

	/*
	   	fold along x=655
	   fold along y=447
	   fold along x=327
	   fold along y=223
	   fold along x=163
	   fold along y=111
	   fold along x=81
	   fold along y=55
	   fold along x=40
	   fold along y=27
	   fold along y=13
	   fold along y=6 */

	foldx(paper, 655)

	fmt.Println(len(paper))

	sheet1 := foldy(paper, 447)
	sheet2 := foldx(sheet1, 327)
	sheet3 := foldy(sheet2, 223)
	sheet4 := foldx(sheet3, 163)
	sheet5 := foldy(sheet4, 111)
	sheet6 := foldx(sheet5, 81)
	sheet7 := foldy(sheet6, 55)
	sheet8 := foldx(sheet7, 40)
	sheet9 := foldy(sheet8, 27)
	sheet10 := foldy(sheet9, 13)
	sheet11 := foldy(sheet10, 6)

	for i := range sheet11 {
		fmt.Println(sheet11[i])
	}

	fmt.Println(counthash(sheet11))

}

func foldy(paper [][]string, line int) [][]string {
	for y := 0; y < line; y++ {
		z := int(math.Abs(float64(y)-float64(line)) + float64(line))
		if z < len(paper) {
			for x := range paper[y] {
				if paper[z][x] == "#" {
					paper[y][x] = "#"
				}
			}
		}
	}

	xerox := paper[0:line]
	return xerox
}

func foldx(paper [][]string, line int) [][]string {
	for x := 0; x < line; x++ {
		z := int(math.Abs(float64(x)-float64(line)) + float64(line))
		for y := range paper {
			if paper[y][z] == "#" {
				paper[y][x] = "#"
			}
		}
	}
	for y := range paper {
		paper[y] = paper[y][:line]
	}
	return paper
}

func counthash(paper [][]string) int {
	count := 0
	for y := range paper {
		for x := range paper[y] {
			if paper[y][x] == "#" {
				count++
			}
		}
	}
	return count
}
