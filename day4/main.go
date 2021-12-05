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

	input, err := i.Strings(2021, 4)
	if err != nil {
		log.Fatal(err)
	}

	//Pull out the Calls in the first line
	var calls []int

	for _, item := range strings.Split(input[0], ",") {
		call, _ := strconv.Atoi(string(item))
		calls = append(calls, call)
	}

	fmt.Println(calls)

	//Make an empty stack of cards
	cardStack := make([][][]int, 0)
	for i := range cardStack {
		cardStack[i] = make([][]int, 5)
		for j := range cardStack[i] {
			cardStack[i][j] = make([]int, 5)
		}
	}

	//fmt.Println(cardStack)

	//tracks what line of each card we're reading in we're on, 0 = the blank line between cards
	cardIndex := 0
	//# of cards
	//cardCount := 0

	//a temp card to fill out and append to our stack
	workingCard := [][]int{
		{5: 0},
		{5: 0},
		{5: 0},
		{5: 0},
		{5: 0},
	}

	//fmt.Println(workingCard)

	//loop through the input starting at line 1
	for i := 1; i < len(input); i++ {
		if cardIndex == 0 {
			// when we see a blank line let's make working card and empty card again
			workingCard = [][]int{
				{5: 0},
				{5: 0},
				{5: 0},
				{5: 0},
				{5: 0},
			}
		}
		if cardIndex%6 != 0 {
			//if it's not the last line, we input it
			line := strings.Fields(input[i])
			intLine := make([]int, 5)
			for i, item := range line {
				intLine[i], _ = strconv.Atoi(string(item))
			}
			//fmt.Println(intLine)
			workingCard[cardIndex-1] = intLine
		}
		cardIndex++
		if cardIndex > 5 {
			cardStack = append(cardStack, workingCard)
			cardIndex = 0
		}
	}

	for _, i := range calls {
		winnerFound := false
		for j := 0; j < len(cardStack); j++ {
			cardStack[j] = markScore(cardStack[j], i)
			winnerFound = checkWin(cardStack[j])
			if winnerFound {
				fmt.Println("Score = ", sumCard(cardStack[j]), "justCalled", i, "Multiplied", i*sumCard(cardStack[j]))
				break
			}
		}
		if winnerFound {
			break
		}
	}
}

func checkWin(card [][]int) bool {
	win := true
	// check for row wins
	for r := 0; r < 5; r++ {
		win = true
		for c := 0; c < 5; c++ {
			if card[r][c] != 0 {
				win = false
			}
			if c == 4 && win {
				return win
			}
		}
	}
	// check for col wins
	for c := 0; c < 5; c++ {
		win = true
		for r := 0; r < 5; r++ {
			if card[r][c] != 0 {
				win = false
			}
			if r == 4 && win {
				return win
			}
		}
	}
	return win
}

func markScore(card [][]int, call int) [][]int {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if card[r][c] == call {
				card[r][c] = 0
				break
			}
		}
	}
	for _, row := range card {
		fmt.Println("call:", call, "\n", row)
	}
	return card
}

func sumCard(card [][]int) int {
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			sum += card[r][c]
		}
	}
	return sum
}
