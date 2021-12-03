package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 3)
	if err != nil {
		log.Fatal(err)
	}

	var gamma = make([]int, len(input[0]))
	var gammaRate = make([]int, len(input[0]))
	var epsilon = make([]int, len(input[0]))
	var epsilonRate = make([]int, len(input[0]))

	for _, byte := range input {
		for j, bit := range byte {
			switch bit {
			case '1':
				gamma[j] += 1
			case '0':
				epsilon[j] += 1
			}
		}
	}

	for i := 0; i < len(input[0]); i++ {
		if gamma[i] > epsilon[i] {
			gammaRate[i] = 1
		} else {
			epsilonRate[i] = 1
		}
	}

	fmt.Println(gamma, "\n", epsilon, "\n", gammaRate, "\n", epsilonRate)
	gammaRateDec, _ := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(gammaRate)), ""), "[]"), 2, 32)
	epsilonRateDec, _ := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(epsilonRate)), ""), "[]"), 2, 32)
	fmt.Println("GammaRate:", gammaRateDec, "EpsilonRate:", epsilonRateDec, "Multiplied:", gammaRateDec*epsilonRateDec)

}
