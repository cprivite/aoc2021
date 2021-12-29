package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 14)
	if err != nil {
		log.Fatal(err)
	}

	var polymer string
	for _, c := range input[0] {
		polymer += string(c)
	}

	pairs := [][]string{}
	for i := 2; i < len(input); i++ {

		s := strings.Split(input[i], " -> ")
		pairs = append(pairs, s)
	}

	paircounts := map[string]int{}

	for i := 0; i < len(polymer)-1; i++ {
		paircounts[string(polymer[i])+string(polymer[i+1])]++
	}

	fmt.Println(paircounts)
	start := time.Now()
	for j := 0; j < 40; j++ {
		newcounts := make(map[string]int)
		fmt.Println("Elapsed: ", time.Since(start))
		fmt.Println("Starting Step:", j)

		for _, s := range pairs {
			if paircounts[s[0]] >= 1 {
				newcounts[string(s[0][0])+s[1]] += paircounts[s[0]]
				newcounts[s[1]+string(s[0][1])] += paircounts[s[0]]
				newcounts[s[0]] -= paircounts[s[0]]
			}
		}
		for k, v := range newcounts {
			paircounts[k] += v
		}
	}
	fmt.Println(paircounts)

	lettercounts := map[string]int{}

	for k, v := range paircounts {
		lettercounts[string(k[0])] += v
		lettercounts[string(k[1])] += v
	}
	fmt.Println(lettercounts)
	for i := range lettercounts {
		lettercounts[i] = 1 + ((lettercounts[i] - 1) / 2)
	}
	fmt.Println(lettercounts)
	/* 	lettercounts["N"]--
	   	lettercounts["B"]-- */

	keys := make([]int, 0, len(lettercounts))

	for _, k := range lettercounts {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	fmt.Println(polymer, lettercounts, keys[0]-keys[len(keys)-1])

}
