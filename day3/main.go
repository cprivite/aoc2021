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

	var ones = make([]int, len(input[0]))
	var gammaRate = make([]string, len(input[0]))
	var zeros = make([]int, len(input[0]))
	var epsilonRate = make([]string, len(input[0]))

	for _, byte := range input {
		for j, bit := range byte {
			switch bit {
			case '1':
				ones[j] += 1
			case '0':
				zeros[j] += 1
			}
		}
	}

	for i := 0; i < len(input[0]); i++ {
		if ones[i] > zeros[i] {
			gammaRate[i] = "1"
			epsilonRate[i] = "0"
		} else {
			epsilonRate[i] = "1"
			gammaRate[i] = "0"
		}
	}

	fmt.Println(" Zeroes", zeros, "\n", "Ones: ", ones, "\n", "GammaRate:  ", gammaRate, "\n", "EpsilonRate:", epsilonRate)
	gammaRateDec, _ := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(gammaRate)), ""), "[]"), 2, 32)
	epsilonRateDec, _ := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(epsilonRate)), ""), "[]"), 2, 32)
	fmt.Println("GammaRate:", gammaRateDec, "EpsilonRate:", epsilonRateDec, "Multiplied:", gammaRateDec*epsilonRateDec)

	//Oxygen Generation
	oxygen := input

	for i := 0; i < 12; i++ {
		var newOx = make([]string, 0)
		zeros := 0
		ones := 0
		var check byte

		for _, byte := range oxygen {
			switch byte[i] {
			case '1':
				ones += 1
			case '0':
				zeros += 1
			}
		}

		if ones >= zeros {
			check = '1'
		} else {
			check = '0'
		}

		for j := 0; j < len(oxygen); j++ {
			if oxygen[j][i] == check {
				newOx = append(newOx, oxygen[j])
			}
		}

		oxygen = newOx
		if len(oxygen) == 1 {
			break
		}
	}

	//CO2 Generation
	co2 := input

	for i := 0; i < 12; i++ {
		var newCo2 = make([]string, 0)
		zeros := 0
		ones := 0
		var keep byte

		for _, byte := range co2 {
			switch byte[i] {
			case '1':
				ones += 1
			case '0':
				zeros += 1
			}
		}

		if ones >= zeros {
			keep = '0'
		} else {
			keep = '1'
		}

		for j := 0; j < len(co2); j++ {
			if co2[j][i] == keep {
				newCo2 = append(newCo2, co2[j])
			}
		}
		co2 = newCo2
		if len(co2) == 1 {
			break
		}
	}

	OxygenRate, _ := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(oxygen)), ""), "[]"), 2, 32)
	fmt.Println("OxygenRate:", OxygenRate)
	CO2Rate, _ := strconv.ParseInt(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(co2)), ""), "[]"), 2, 32)
	fmt.Println("CO2Rate:", CO2Rate)
	fmt.Println("Multiplied:", OxygenRate*CO2Rate)

}
