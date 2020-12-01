package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2020/day1/input.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)
	}
	var nums []int
	scanner1 := bufio.NewScanner(file)
	for scanner1.Scan() {
		lineStr1 := scanner1.Text()
		num, _ := strconv.Atoi(lineStr1)
		nums = append(nums, num)
	}
	for _, num1 := range nums {
		for _, num2 := range nums {
			for _, num3 := range nums {
				if num1+num2+num3 == 2020 {
					fmt.Println(num1 * num2 * num3)
					break
				}
			}
		}
	}
}
