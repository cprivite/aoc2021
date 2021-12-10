package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 9)
	if err != nil {
		log.Fatal(err)
	}

	mappy := [][]int{}
	coords := [][2]int{}
	sum := 0

	for _, line := range input {
		intArray := []int{}
		for _, char := range line {
			i, _ := strconv.Atoi(string(char))
			intArray = append(intArray, i)
		}
		mappy = append(mappy, intArray)
	}

	mappy2 := make([][]int, len(mappy))
	for i := range mappy2 {
		mappy2[i] = make([]int, len(mappy[0]))
	}

	for i := 0; i < len(mappy); i++ {
		for j := 0; j < len(mappy[i]); j++ {
			c := mappy[i][j]
			if i == 0 && j == 0 {
				if mappy[i+1][j] > c && mappy[i][j+1] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if i == 0 && j < len(mappy[i])-1 {
				if mappy[i][j-1] > c && mappy[i+1][j] > c && mappy[i][j+1] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if i == 0 && j == len(mappy[i])-1 {
				if mappy[i][j-1] > c && mappy[i+1][j] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if i < len(mappy)-1 && j == 0 {
				if mappy[i-1][j] > c && mappy[i+1][j] > c && mappy[i][j+1] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if i < len(mappy)-1 && j == len(mappy[i])-1 {
				if mappy[i-1][j] > c && mappy[i][j-1] > c && mappy[i+1][j] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if i == len(mappy)-1 && j == 0 {
				if mappy[i-1][j] > c && mappy[i][j+1] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if i == len(mappy)-1 && j < len(mappy[i])-1 {
				if mappy[i-1][j] > c && mappy[i][j-1] > c && mappy[i][j+1] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if i == len(mappy)-1 && j == len(mappy[i])-1 {
				if mappy[i-1][j] > c && mappy[i][j-1] > c {
					coords = append(coords, [2]int{i, j})
					sum += mappy[i][j] + 1
				}
			} else if mappy[i-1][j] > c && mappy[i][j-1] > c && mappy[i+1][j] > c && mappy[i][j+1] > c {
				coords = append(coords, [2]int{i, j})
				sum += mappy[i][j] + 1
			}
		}
	}

	basinSizes := []int{}

	for _, coord := range coords {
		basinSize := 0
		i := coord[0]
		j := coord[1]

		queue := [][2]int{}
		queue = append(queue, [2]int{i, j})

		for len(queue) > 0 {
			i := queue[0][0]
			j := queue[0][1]
			queue = queue[1:]
			if mappy2[i][j] != 9 {
				basinSize++
				mappy2[i][j] = 9
			}
			if j > 0 && mappy[i][j-1] != 9 && mappy2[i][j-1] != 9 {
				queue = append(queue, [2]int{i, j - 1})
			}
			if j < len(mappy[i])-1 && mappy[i][j+1] != 9 && mappy2[i][j+1] != 9 {
				queue = append(queue, [2]int{i, j + 1})
			}
			if i > 0 && mappy[i-1][j] != 9 && mappy2[i-1][j] != 9 {
				queue = append(queue, [2]int{i - 1, j})
			}
			if i < len(mappy)-1 && mappy[i+1][j] != 9 && mappy2[i+1][j] != 9 {
				queue = append(queue, [2]int{i + 1, j})
			}
		}
		basinSizes = append(basinSizes, basinSize)
	}
	sort.Ints(basinSizes)
	fmt.Println(basinSizes)
	fmt.Println(basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3])
}
