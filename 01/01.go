package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type TwoSumTo2020 struct {
	a int
	b int
}

type ThreeSumTo2020 struct {
	a int
	b int
	c int
}

func FindTwoSumTo2020(values []int) TwoSumTo2020 {
	for i := range values {
		for j := range values {
			if j == i {
				continue
			}
			if values[i]+values[j] == 2020 {
				return TwoSumTo2020{values[i], values[j]}
			}
		}
	}
	return TwoSumTo2020{0, 0}
}

func FindThreeSumTo2020(values []int) ThreeSumTo2020 {
	for i := range values {
		for j := range values {
			if i == j {
				continue
			}
			for k := range values {
				if i == j || j == k {
					continue
				}
				if values[i]+values[j]+values[k] == 2020 {
					return ThreeSumTo2020{values[i], values[j], values[k]}
				}
			}
		}
	}
	return ThreeSumTo2020{0, 0, 0}
}

func main() {
	data, err := ioutil.ReadFile("01_input.txt")
	if err != nil {
		panic(err)
	}

	txt := string(data)
	lines := strings.Split(txt, "\n")
	var values []int

	for _, i := range lines {
		j, _ := strconv.Atoi(i)
		values = append(values, j)
	}

	twoSumTo2020 := FindTwoSumTo2020(values)

	fmt.Println("Pt1 Sums")
	fmt.Println(twoSumTo2020)
	fmt.Println("Pt1 Multiplication")
	fmt.Println(twoSumTo2020.a * twoSumTo2020.b)

	threeSumTo2020 := FindThreeSumTo2020(values)

	fmt.Println("Pt2 Sums")
	fmt.Println(threeSumTo2020)
	fmt.Println("Pt2 Multiplication")
	fmt.Println(threeSumTo2020.a * threeSumTo2020.b * threeSumTo2020.c)
}
