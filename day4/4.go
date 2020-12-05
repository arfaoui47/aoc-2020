package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	file, err := os.Open("2020/day4/input.txt")
	//file, err := os.Open("2020/day4/input_small.txt")
	if err != nil {
		log.Fatalf("shiit: %s", err)
	}

	scanner := bufio.NewScanner(file)
	var newPassenger string
	valid := 0
	for scanner.Scan() {
		line := scanner.Text()
		newPassenger += " " + line
		newPassenger = strings.ReplaceAll(newPassenger, "\n", " ")
		if line == "" {
			newPassenger = strings.Trim(newPassenger, " ")
			items := strings.Split(newPassenger, " ")
			var ids []string
			var values []string
			for _, item := range items {
				itemS := strings.Split(item, ":")
				ids = append(ids, itemS[0])
				values = append(values, itemS[1])
			}
			if AllItemsIn(requiredFields, ids) {
				if checkValidValues(ids, values) {
					valid++
					fmt.Println("passenger", newPassenger)

				}else {
					//fmt.Println("passenger", newPassenger)
				}
			}
			newPassenger = ""
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

func checkValidValues(ids, values []string) bool {
	for index, id := range ids {
		if id == "byr" {
			value := values[index]
			valueInt, _ := strconv.Atoi(value)
			if !(1920 <= valueInt && valueInt <= 2002) {
				fmt.Println("byr: ", value)
				return false
			}
		}
		if id == "iyr" {
			value := values[index]
			valueInt, _ := strconv.Atoi(value)
			if !(2010 <= valueInt && valueInt <= 2020) {
				fmt.Println("iyr: ", value)
				return false
			}
		}
		if id == "eyr" {
			value := values[index]
			valueInt, _ := strconv.Atoi(value)
			if !(2020 <= valueInt && valueInt <= 2030) {
				fmt.Println("eyr: ", value)
				return false
			}
		}
		if id == "hgt" {
			value := values[index]
			if value[len(value)-2:] == "in" {
				valueInt, _ := strconv.Atoi(value[:len(value)-2])
				if !(59 <= valueInt && valueInt <= 76) {
					fmt.Println("hgt in: ", value)
					return false
				}
			}
			if value[len(value)-2:] == "cm" {
				valueInt, _ := strconv.Atoi(value[:len(value)-2])
				if !(150 <= valueInt && valueInt <= 193) {
					fmt.Println("hgt cm: ", value)
					return false
				}
			}
		}
		if id == "hcl" {
			value := values[index]
			if !(value[0] == '#') {
				fmt.Println("hcl #: ", value)
				return false
			}
			re := regexp.MustCompile("^[a-z0-9]*$")
			if !(re.MatchString(value[1:])) {
				fmt.Println("hcl alphanum: ", value)
				return false
			}
		}
		if id == "ecl" {
			eyes := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			value := values[index]
			if !Contain(eyes, value) {
				fmt.Println("ecl: ", value)
				return false
			}
		}
		if id == "pid" {
			value := values[index]
			if !(len(value) == 9) {
				fmt.Println("pid: len", value)
				return false
			}
			if _, err := strconv.Atoi(value); err != nil {
				fmt.Println("pid: int", value)
				return false
			}
		}
	}
	return true
}
