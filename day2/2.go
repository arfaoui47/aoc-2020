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
	file, err := os.Open("2020/day2/input.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)
	}
	validCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		rangeChar := split[0]
		str := split[1]
		rangeCharMap := strings.Split(rangeChar, " ")
		range_ := rangeCharMap[0]
		char := rangeCharMap[1]
		bounds := strings.Split(range_, "-")
		lbound := bounds[0]
		highbound := bounds[1]
		charCount := strings.Count(str, char)
		lbound_, _ := strconv.Atoi(lbound)
		highbound_, _ := strconv.Atoi(highbound)
		if lbound_ <= charCount && charCount <=highbound_ {
			validCount++
		}
	}
	fmt.Println(validCount)
}
