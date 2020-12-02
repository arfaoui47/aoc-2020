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
		str = strings.TrimSpace(str)
		rangeCharMap := strings.Split(rangeChar, " ")
		range_ := rangeCharMap[0]
		char := rangeCharMap[1]
		indexes := strings.Split(range_, "-")
		lindex := indexes[0]
		highindex := indexes[1]
		lindex_, _ := strconv.Atoi(lindex)
		highindex_, _ := strconv.Atoi(highindex)

		a := char == string(str[lindex_ - 1])
		b := char == string(str[highindex_ - 1])

		if  (a || b) && !(a && b){
			validCount++
		}
	}
	fmt.Println(validCount)
}
