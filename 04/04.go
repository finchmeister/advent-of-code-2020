package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func FindNoOfValidPassportsPt1(passportRawData []string) int {
	return validatePassports(passportRawData, validatorPt1)
}

func FindNoOfValidPassportsPt2(passportRawData []string) int {
	return validatePassports(passportRawData, validatorPt2)
}

func validatePassports(passportRawData []string, validator func(passportRawRow string) bool) int {
	i := 0
	for _, passportRawRow := range passportRawData {
		if validator(passportRawRow) {
			i++
		}
	}

	return i
}

func validatorPt1(passportRawRow string) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, requiredField := range requiredFields {
		matched, _ := regexp.Match(requiredField+":", []byte(passportRawRow))
		if matched == false {
			return false
		}
	}

	return true
}

func validatorPt2(passportRawRow string) bool {
	if isByrValid(getValue("byr", passportRawRow)) == false {
		return false
	}

	if isIyrValid(getValue("iyr", passportRawRow)) == false {
		return false
	}

	if isEyrValid(getValue("eyr", passportRawRow)) == false {
		return false
	}

	if isHgtValid(getValue("hgt", passportRawRow)) == false {
		return false
	}

	if isHclValid(getValue("hcl", passportRawRow)) == false {
		return false
	}

	if isEclValid(getValue("ecl", passportRawRow)) == false {
		return false
	}

	if isPidValid(getValue("pid", passportRawRow)) == false {
		return false
	}

	return true
}

func getValue(key string, rawData string) string {
	re := regexp.MustCompile(key + ":([0-9a-z#]+)")
	matches := re.FindStringSubmatch(rawData)
	if len(matches) == 0 {
		return ""
	}

	return matches[1]
}

func isByrValid(byr string) bool {
	return validateDate(1920, 2002, byr)
}

func isIyrValid(iyr string) bool {
	return validateDate(2010, 2020, iyr)
}

func isEyrValid(Eyr string) bool {
	return validateDate(2020, 2030, Eyr)
}

func isHgtValid(hgt string) bool {
	re := regexp.MustCompile(`^(\d{3})cm$`)
	matches := re.FindStringSubmatch(hgt)
	if len(matches) > 0 {
		hgtInCm, _ := strconv.Atoi(matches[1])
		if hgtInCm > 193 {
			return false
		}
		if hgtInCm < 150 {
			return false
		}

		return true
	}

	re = regexp.MustCompile(`^(\d{2})in$`)
	matches = re.FindStringSubmatch(hgt)
	if len(matches) > 0 {
		hgtInCm, _ := strconv.Atoi(matches[1])
		if hgtInCm > 76 {
			return false
		}
		if hgtInCm < 59 {
			return false
		}

		return true
	}

	return false
}

func isHclValid(hcl string) bool {
	return isRegexMatch(`^#[0-9a-f]{6}$`, hcl)
}

func isEclValid(hcl string) bool {
	return isRegexMatch(`^amb|blu|brn|gry|grn|hzl|oth$`, hcl)
}

func isPidValid(pid string) bool {
	return isRegexMatch(`^[0-9]{9}$`, pid)
}

func isRegexMatch(pattern string, s string) bool {
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func validateDate(min int, max int, subject string) bool {
	match, _ := regexp.MatchString(`^\d{4}$`, subject)
	if match == false {
		return false
	}

	value, _ := strconv.Atoi(subject)

	if value < min {
		return false
	}

	if value > max {
		return false
	}

	return true
}

func parseBatchFile(batchFile string) []string {
	return strings.Split(batchFile, "\n\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("04_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1 No of Valid Passports")
	fmt.Println(FindNoOfValidPassportsPt1(parseBatchFile(loadFile())))
	fmt.Println("Pt2 No of Valid Passports")
	fmt.Println(FindNoOfValidPassportsPt2(parseBatchFile(loadFile())))
}
