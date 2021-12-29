package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/RyanCarrier/dijkstra"
	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 15)
	if err != nil {
		log.Fatal(err)
	}

	cavern := make([][]int, 0, 5*len(input))

	for i := 0; i < 5; i++ {
		for _, line := range input {
			lineInts := make([]int, 0)
			for j := 0; j < 5; j++ {
				for _, strInt := range line {
					innt, _ := strconv.Atoi(string(strInt))
					if innt+j+i <= 9 {
						lineInts = append(lineInts, innt+j+i)
					} else {
						lineInts = append(lineInts, innt+j+i-9)
					}
				}
			}
			cavern = append(cavern, lineInts)
		}
	}

	for i := range cavern {
		fmt.Println(cavern[i])
	}
	graph := dijkstra.NewGraph()

	//enumerate the # of elements in the slice and add a vertex for each
	for i, row := range cavern {
		for j := range row {
			graph.AddVertex(i*len(row) + j)
		}
	}

	for i, row := range cavern {
		for j, value := range row {
			index := i*len(row) + j
			if i > 0 {
				graph.AddArc((i-1)*len(row)+j, index, int64(value))
			}
			if i < len(cavern)-1 {
				graph.AddArc((i+1)*len(row)+j, index, int64(value))
			}
			if j > 0 {
				graph.AddArc((i)*len(row)+j-1, index, int64(value))
			}
			if j < len(row)-1 {
				graph.AddArc((i)*len(row)+j+1, index, int64(value))
			}
		}
	}

	best, _ := graph.Shortest(0, len(cavern)*len(cavern[0])-1)

	fmt.Println(best.Distance)

}
