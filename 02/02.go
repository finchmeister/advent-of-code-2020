package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type PasswordPolicyPassword struct {
	a        int
	b        int
	c        string
	password string
}

func FindValidPasswordsSchemeACount(passwordPolicyPasswords []PasswordPolicyPassword) int {
	i := 0
	for _, passwordPolicyPassword := range passwordPolicyPasswords {
		if isPasswordSchemeAValid(passwordPolicyPassword) {
			i++
		}
	}

	return i
}

func isPasswordSchemeAValid(passwordPolicyPassword PasswordPolicyPassword) bool {
	regex := regexp.MustCompile(passwordPolicyPassword.c)
	matches := regex.FindAllStringIndex(passwordPolicyPassword.password, -1)

	return passwordPolicyPassword.a <= len(matches) && passwordPolicyPassword.b >= len(matches)
}

func FindValidPasswordsSchemeBCount(passwordPolicyPasswords []PasswordPolicyPassword) int {
	i := 0
	for _, passwordPolicyPassword := range passwordPolicyPasswords {
		if isPasswordSchemeBValid(passwordPolicyPassword) {
			i++
		}
	}

	return i
}

func isPasswordSchemeBValid(passwordPolicyPassword PasswordPolicyPassword) bool {
	posAMatch := string(passwordPolicyPassword.password[passwordPolicyPassword.a-1]) == passwordPolicyPassword.c
	posBMatch := string(passwordPolicyPassword.password[passwordPolicyPassword.b-1]) == passwordPolicyPassword.c

	if posAMatch && posBMatch {
		return false
	}

	if !posAMatch && !posBMatch {
		return false
	}

	return true
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
	fmt.Println("Scheme A")
	fmt.Println(FindValidPasswordsSchemeACount(passwordPolicyPasswords))
	fmt.Println("Scheme B")
	fmt.Println(FindValidPasswordsSchemeBCount(passwordPolicyPasswords))
}
