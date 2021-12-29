package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Bytes(2021, 16)
	if err != nil {
		log.Fatal(err)
	}

	var binary string
	for _, b := range input {
		i, _ := strconv.ParseUint(string(b), 16, 4)
		bi := fmt.Sprintf("%04b", i)
		binary += bi
	}

	fmt.Println(binary)
	//packet with subpackets
	//0 - version
	//1 - version
	//2 - version
	//3 - type id
	//4 - type id
	//5 - type id
	//6 - length type ID || continue/no contine on literal
	//7-21 literal or subpacket field
	//18-21 literal continued or the subpackets

	// type table
	// 4 = literal value, single binary number with leading zero padding to multiple of 4 bits broken up and padded with a 1 unless it's the last part of the number, then they pad with a 0.
	// !4 = operator packet, followed by length type bit

	//lenght type ID
	//0 - subpacket length is next 15 bits
	//1 - subpacket count is next 11 bits

	parseThis(binary)
}
func parseThis(binary string) string {

	return ""
}
