package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//file, err := os.Open("2020/day3/input_small.txt")
	file, err := os.Open("2020/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	xslope_list := []int{1, 3, 5, 7, 1}
	yslope_list := []int{1, 1, 1, 1, 2}
	var data [][]string
	for scanner.Scan() {
		line := scanner.Text()
		line_b := strings.Repeat(line, 100)
		data = append(data, strings.Split(line_b, ""))

	}
	var record []int
	prod:= 1
	for idx, iter := range xslope_list {
		slope := 0
		treeCount := 0

		for row_count, row := range data {
			if row_count%yslope_list[idx] != 0 {
				continue
			}
			if row[slope] == "#" {
				treeCount++
			}
			slope += iter
		}
		prod*= treeCount
		record = append(record, treeCount)

	}
	fmt.Println(record)
	fmt.Println(prod)
}
