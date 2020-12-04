package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	//required_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

	file, err := os.Open("2020/day4/input.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)
	}

	scanner := bufio.NewScanner(file)
	var new_passenger string
	valid := 0
	for scanner.Scan() {
		line := scanner.Text()
		new_passenger += " " + line
		new_passenger = strings.ReplaceAll(new_passenger, "\n", " ")
		if line == "" {
			new_passenger = strings.Trim(new_passenger, " ")
			items := strings.Split(new_passenger, " ")
			var ids []string
			for _, item := range items {
				itemS := strings.Split(item, ":")
				ids = append(ids, itemS[0])
			}
			new_passenger = ""
			if AllItemsIn(required_fields, ids) {
				valid++
			}
		}
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
