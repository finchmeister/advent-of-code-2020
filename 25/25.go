package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const modulo = 20201227

func FindEncryptionKeyPt1(publicKeyA, publicKeyB int) int {
	subjectNumber := 7
	publicKeyALoopSize := getLoopSize(subjectNumber, publicKeyA)

	return applyLoop(publicKeyB, publicKeyALoopSize)
}

func getLoopSize(subjectNumber, publicKey int) int {
	result := 1
	i := 0
	for result != publicKey {
		result = mod(result*subjectNumber, modulo)
		i++
	}

	return i
}

func applyLoop(subjectNumber, loop int) int {
	result := 1
	for i := 0; i < loop; i++ {
		result = mod(result*subjectNumber, modulo)
	}

	return result
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func parse(input string) (int, int) {
	rawKeys := strings.Split(input, "\n")

	publicKeyA, _ := strconv.Atoi(rawKeys[0])
	publicKeyB, _ := strconv.Atoi(rawKeys[1])

	return publicKeyA, publicKeyB
}

func loadFile() string {
	data, err := ioutil.ReadFile("25_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindEncryptionKeyPt1(parse(loadFile())))
}
