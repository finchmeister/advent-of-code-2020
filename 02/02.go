package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type PasswordPolicyPassword struct {
	min      int
	max      int
	c        string
	password string
}

func FindValidPasswordsCount(passwordPolicyPasswords []PasswordPolicyPassword) int {
	i := 0
	for _, passwordPolicyPassword := range passwordPolicyPasswords {
		if isPasswordValid(passwordPolicyPassword) {
			i++
		}
	}

	return i
}

func isPasswordValid(passwordPolicyPassword PasswordPolicyPassword) bool {
	regex := regexp.MustCompile(passwordPolicyPassword.c)
	matches := regex.FindAllStringIndex(passwordPolicyPassword.password, -1)

	return passwordPolicyPassword.min <= len(matches) && passwordPolicyPassword.max >= len(matches)
}

func loadFile() []PasswordPolicyPassword {
	data, err := ioutil.ReadFile("02_input.txt")
	if err != nil {
		panic(err)
	}

	txt := string(data)
	lines := strings.Split(txt, "\n")

	var passwordPolicyPasswords []PasswordPolicyPassword

	for _, line := range lines {
		passwordPolicyPasswords = append(passwordPolicyPasswords, createPasswordPolicyPasswordFromString(line))
	}

	return passwordPolicyPasswords
}

func createPasswordPolicyPasswordFromString(s string) PasswordPolicyPassword {
	r := regexp.MustCompile(`(\d+)-(\d+) (\D): (\D+)`)
	m := r.FindSubmatch([]byte(s))

	min, _ := strconv.Atoi(string(m[1]))
	max, _ := strconv.Atoi(string(m[2]))

	return PasswordPolicyPassword{
		min,
		max,
		string(m[3]),
		string(m[4]),
	}
}

func main() {
	passwordPolicyPasswords := loadFile()
	fmt.Println(FindValidPasswordsCount(passwordPolicyPasswords))
}
