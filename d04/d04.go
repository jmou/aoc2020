package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("d04/input")
	credentials := strings.Split(string(fileBytes), "\n\n")

	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	count_p1 := 0
	count_p2 := 0
	for _, credential := range credentials {
		fields := make(map[string]string)
		for _, field := range strings.Fields(credential) {
			fields[field[0:3]] = field[4:]
		}

		valid := true
		for _, field := range requiredFields {
			if _, ok := fields[field]; !ok {
				valid = false
				break
			}
		}
		if valid {
			count_p1++
		}

		if x := fields["byr"]; !(len(x) == 4 && x >= "1920" && x <= "2002") {
			continue
		}
		if x := fields["iyr"]; !(len(x) == 4 && x >= "2010" && x <= "2020") {
			continue
		}
		if x := fields["eyr"]; !(len(x) == 4 && x >= "2020" && x <= "2030") {
			continue
		}

		hgt := fields["hgt"]
		hgt_val, hgt_unit := hgt[:len(hgt)-2], hgt[len(hgt)-2:]
		if hgt_unit == "cm" {
			if !(len(hgt_val) == 3 && hgt_val >= "150" && hgt_val <= "193") {
				continue
			}
		} else if hgt_unit == "in" {
			if !(len(hgt_val) == 2 && hgt_val >= "59" && hgt_val <= "76") {
				continue
			}
		} else {
			continue
		}

		if matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, fields["hcl"]); !matched {
			continue
		}
		if x := fields["ecl"]; x != "amb" && x != "blu" && x != "brn" &&
			x != "gry" && x != "grn" && x != "hzl" && x != "oth" {
			continue
		}
		if matched, _ := regexp.MatchString(`^[0-9]{9}$`, fields["pid"]); !matched {
			continue
		}
		count_p2++
	}

	fmt.Println(count_p1)
	fmt.Println(count_p2)
}
