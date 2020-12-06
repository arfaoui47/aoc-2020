package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//file, err := os.Open("2020/day6/input_small.txt")
	file, err := os.Open("2020/day6/input.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)
	}

	scanner := bufio.NewScanner(file)
	var newGroup string
	valid := 0
	for scanner.Scan() {
		line := scanner.Text()
		newGroup += line
		newGroup = strings.ReplaceAll(newGroup, "\n", " ")
		set := make(map[string]bool)
		if line == "" {
			items := strings.Split(newGroup, "")
			for _, item := range items {
				set[item] = true
			}
			newGroup = ""
		}
		fmt.Println("set", set)
		valid += len(set)
	}

	fmt.Println("Valid passports: ", valid)
}

func AllItemsIn(l1, l2 []string) bool {
	for _, item := range l1 {
		if !Contain(l2, item) {
			return false
		}
	}
	return true
}

func Contain(list []string, item string) bool {
	for _, it := range list {
		if item == it {
			return true
		}
	}
	return false
}
