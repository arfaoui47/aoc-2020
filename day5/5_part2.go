package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file, err := os.Open("2020/day5/input_small.txt")
	file, err := os.Open("2020/day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var seats []int
	var max int
	var diff []int

	for scanner.Scan() {
		line := scanner.Text()
		firsPart := line[:len(line)-3]
		secondPart := line[len(line)-3:]
		firsPartBinary := intToBinary(strings.ReplaceAll(strings.ReplaceAll(firsPart, "B", "1"), "F", "0"))
		secondPartBinary := intToBinary(strings.ReplaceAll(strings.ReplaceAll(secondPart, "R", "1"), "L", "0"))
		if int(firsPartBinary*8+secondPartBinary) > max {
			max = int(firsPartBinary*8 + secondPartBinary)
		}
		seats = append(seats, int(firsPartBinary*8+secondPartBinary))

	}
	mapingSeats := make(map[int]bool)
	var allSeats []int
	for i := 1; i <= max; i++ {
		allSeats = append(allSeats, i)
	}
	for _, seat := range seats {
		mapingSeats[seat] = true
	}
	for _, seat := range allSeats {
		if _, ok := mapingSeats[seat]; !ok {
			diff = append(diff, seat)
		}
	}

	fmt.Println("missing seat", diff[len(diff)-1])
}

func intToBinary(part string) int64 {
	integer, _ := strconv.ParseInt(part, 2, 64)
	return integer
}
