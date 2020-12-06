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
	var max int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		firsPart := line[:len(line)-3]
		secondPart := line[len(line)-3:]
		firsPartBinary := IntToBinary(strings.ReplaceAll(strings.ReplaceAll(firsPart, "B", "1"), "F", "0"))
		secondPartBinary := IntToBinary(strings.ReplaceAll(strings.ReplaceAll(secondPart, "R", "1"), "L", "0"))
		if firsPartBinary*8+secondPartBinary > max {
			max = firsPartBinary*8 + secondPartBinary
		}
	}
	fmt.Println("Max sit", max)

}

func IntToBinary(part string) int64 {
	integer, _ := strconv.ParseInt(part, 2, 64)
	return integer
}
