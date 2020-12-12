package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./4/input.txt")
	passports := make([][]string, 0)
	var passport []string
	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = []string{}
			continue
		}

		passport = append(passport, strings.Split(line, " ")...)
	}
	passports = append(passports, passport)

	valid := 0
	for _, pairs := range passports {
		passport := unpack(pairs)
		if validate(passport) {
			valid++
		}
	}

	log.Printf("%d passports found\n", len(passports))
	log.Printf("%d valid passports found\n", valid)
}

func partTwo() {
	lines := aoc.ReadInput("./4/input.txt")
	passports := make([][]string, 0)
	var passport []string
	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = []string{}
			continue
		}

		passport = append(passport, strings.Split(line, " ")...)
	}
	passports = append(passports, passport)

	valid := 0
	for _, pairs := range passports {
		passport := unpack(pairs)
		if validateStrict(passport) {
			valid++
		}
	}

	log.Printf("%d passports found\n", len(passports))
	log.Printf("%d valid passports found\n", valid)
}

type ValidationFunc func(string) bool

var (
	BirthYear      = "byr"
	IssueYear      = "iyr"
	ExpirationYear = "eyr"
	Height         = "hgt"
	HairColour     = "hcl"
	EyeColour      = "ecl"
	PassportId     = "pid"
	CountryId      = "cid"
	requiredFields = map[string]ValidationFunc{
		BirthYear: func(s string) bool {
			if len(s) != 4 {
				return false
			}

			i, err := strconv.Atoi(s)
			if err != nil {
				return false
			}

			return i >= 1920 && i <= 2002
		},
		IssueYear: func(s string) bool {
			if len(s) != 4 {
				return false
			}

			i, err := strconv.Atoi(s)
			if err != nil {
				return false
			}

			return i >= 2010 && i <= 2020
		},
		ExpirationYear: func(s string) bool {
			if len(s) != 4 {
				return false
			}

			i, err := strconv.Atoi(s)
			if err != nil {
				return false
			}

			return i >= 2020 && i <= 2030
		},
		Height: func(s string) bool {
			pattern := regexp.MustCompile(`^(\d+)(cm|in)$`)
			if !pattern.MatchString(s) {
				return false
			}

			match := pattern.FindStringSubmatch(s)
			if len(match) != 3 {
				return false
			}
			i, err := strconv.Atoi(match[1])
			if err != nil {
				return false
			}

			switch match[2] {
			case "cm":
				return i >= 150 && i <= 193
			case "in":
				return i >= 59 && i <= 76
			default:
				return false
			}
		},
		HairColour: func(s string) bool {
			return regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(s)
		},
		EyeColour: func(s string) bool {
			colours := []string{"amb","blu","brn","gry","grn","hzl","oth"}
			for _, colour := range colours {
				if colour == s {
					return true
				}
			}

			return false
		},
		PassportId: func(s string) bool {
			matched := regexp.MustCompile(`^\d{9}$`).MatchString(s)
			return matched
		},
	}
)

func unpack(pairs []string) map[string]string {
	fields := make(map[string]string)
	for _, raw := range pairs {
		pair := strings.Split(raw, ":")
		fields[pair[0]] = pair[1]
	}

	return fields
}

func validate(fields map[string]string) bool {
	for requiredField := range requiredFields {
		if _, ok := fields[requiredField]; !ok {
			return false
		}
	}

	return true
}

func validateStrict(fields map[string]string) bool {
	for requiredField, validate := range requiredFields {
		val, ok := fields[requiredField]
		if !ok {
			return false
		}

		if !validate(val) {
			return false
		}
	}

	return true
}
