package main

import (
	"fmt"
	"log"
	"sort"
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

	input, err := i.Strings(2021, 8)
	if err != nil {
		log.Fatal(err)
	}

	sequences := [][]string{}
	outputs := [][]string{}

	for _, line := range input {
		splitString := strings.Fields(line)
		for i, s := range splitString {
			splitString[i] = normalize(s)
		}
		sequences = append(sequences, splitString[:10])
		outputs = append(outputs, splitString[11:])
	}

	//Part 1
	for i := 0; i < len(sequences); i++ {
		fmt.Println("Sequences", sequences[i], "Outputs", outputs[i])
	}
	fmt.Println(countUniques(outputs))

	//Part 2
	count := 0
	for i := 0; i < len(outputs); i++ {
		myMap := decodeSequences(sequences[i])
		var s string
		for _, j := range outputs[i] {
			s += strconv.Itoa(myMap[j])
		}
		c, _ := strconv.Atoi(s)
		count += c
	}
	fmt.Println(count)
}

func countUniques(sliceSliceStrings [][]string) int {
	count := 0
	for _, sliceStrings := range sliceSliceStrings {
		for _, s := range sliceStrings {
			if len(s) != 5 && len(s) != 6 {
				count++
			}
		}
	}
	return count
}

func decodeSequences(sequences []string) map[string]int {
	var one, four, seven, eight string
	var five, six []string

	for _, s := range sequences {
		switch len(s) {
		case 2:
			one = s
		case 3:
			seven = s
		case 4:
			four = s
		case 7:
			eight = s
		case 5:
			five = append(five, s)
		case 6:
			six = append(six, s)
		}
	}

	fmt.Println(one, "", four, "", seven, "", eight, "", five, "", six)

	out := map[string]int{
		one:   1,
		four:  4,
		seven: 7,
		eight: 8,
	}

	var sixString string

	//figure out 9, 0, 6
	for _, s := range six {
		if strings.Contains(s, string(four[0])) && strings.Contains(s, string(four[1])) && strings.Contains(s, string(four[2])) && strings.Contains(s, string(four[3])) {
			out[s] = 9
		} else if strings.Contains(s, string(one[0])) && strings.Contains(s, string(one[1])) {
			out[s] = 0
		} else {
			out[s] = 6
			sixString = s
		}

	}

	//figure out 2, 3, 5
	for _, s := range five {
		if strings.Contains(s, string(seven[0])) && strings.Contains(s, string(seven[1])) && strings.Contains(s, string(seven[2])) {
			out[s] = 3
		} else if strings.Contains(sixString, string(s[0])) && strings.Contains(sixString, string(s[1])) && strings.Contains(sixString, string(s[2])) && strings.Contains(sixString, string(s[3])) && strings.Contains(sixString, string(s[4])) {
			out[s] = 5
		} else {
			out[s] = 2
		}
	}

	return out
}

func normalize(s string) string {
	arr := strings.Split(s, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}
