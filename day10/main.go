package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 10)
	if err != nil {
		log.Fatal(err)
	}
	incQueue := [][]rune{}
	totalScore := 0
	for _, s := range input {
		queue := []rune{}
		corrupt := false
		score := 0
		for _, t := range s {
			switch t {
			case '{':
				queue = append(queue, t)
			case '[':
				queue = append(queue, t)
			case '(':
				queue = append(queue, t)
			case '<':
				queue = append(queue, t)
			case '}':
				if queue[len(queue)-1] == '{' {
					queue = queue[:len(queue)-1]
				} else {
					if !corrupt {
						score += 1197
						corrupt = true
					}
				}
			case ']':
				if queue[len(queue)-1] == '[' {
					queue = queue[:len(queue)-1]
				} else {
					if !corrupt {
						score += 57
						corrupt = true
					}
				}
			case ')':
				if queue[len(queue)-1] == '(' {
					queue = queue[:len(queue)-1]
				} else {
					if !corrupt {
						score += 3
						corrupt = true
					}
				}
			case '>':
				if queue[len(queue)-1] == '<' {
					queue = queue[:len(queue)-1]
				} else {
					if !corrupt {
						score += 25137
						corrupt = true
					}
				}
			}
		}
		fmt.Println("Corrupt?", corrupt, "Score:", score, "Queue", string(queue))
		totalScore += score
		if !corrupt {
			incQueue = append(incQueue, queue)
		}
	}

	//part 1
	fmt.Println("Total Score:", totalScore)

	autoScore := []int{}
	for _, s := range incQueue {
		queue := []rune{}
		score := 0
		for i := len(s) - 1; i >= 0; i-- {
			switch s[i] {
			case '{':
				//queue = append(queue, t)
				score = score*5 + 3
			case '[':
				//queue = append(queue, t)
				score = score*5 + 2
			case '(':
				//queue = append(queue, t)
				score = score*5 + 1
			case '<':
				//queue = append(queue, t)
				score = score*5 + 4
			}

		}
		fmt.Println("Score:", score, "Queue", string(queue))
		autoScore = append(autoScore, score)
	}
	sort.Ints(autoScore)
	midDex := int(len(autoScore) / 2)
	fmt.Println("TotalScore:", autoScore[midDex])
}
