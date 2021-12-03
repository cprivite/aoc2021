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

	input, err := i.Strings(2021, 2)
	if err != nil {
		log.Fatal(err)
	}

	x := 0
	y := 0
	z := 0

	for i, text := range input {
		vector := strings.Split(text, " ")
		direction := vector[0]
		distance, _ := strconv.Atoi(vector[1])
		fmt.Println("index:", i, "\tDirection:", direction, "\tValue:", distance)
		switch direction {
		case "forward":
			x = x + distance
			y = y + (z * distance)
		case "up":
			z = z - distance
		case "down":
			z = z + distance
		}
		fmt.Println("X:", x, "Y:", y, "Z:", z)
	}

	fmt.Println("X:", x, "Y:", y, "Z:", z, "Multiplied", x*y)
}
