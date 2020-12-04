package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2020/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	slope := 0
	treeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		line_b := strings.Repeat(line, 32)
		if line_b[slope] == '#' {
			treeCount++
		}
		slope += 3
	}
	fmt.Println(treeCount)
}
