package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func FindNoOfValidPassportsPt1(passportRawData []string) int {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	i := 0

	for _, passportRawRow := range passportRawData {
		if doesPassportRawRowContainRequiredFields(passportRawRow, requiredFields) {
			i++
		}
	}

	return i
}

func doesPassportRawRowContainRequiredFields(passportRawRow string, requiredFields []string) bool {
	for _, requiredField := range requiredFields {
		matched, _ := regexp.Match(requiredField+":", []byte(passportRawRow))
		if matched == false {
			return false
		}
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
}
