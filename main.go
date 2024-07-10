package main

import (
	"fmt"
	"strconv"
)

func IsValid(id string) bool {
	sum := 0
	check := 0
	last := 0

	if len(id) != 13 {
		return false
	}
	index := 0

	for i := 13; i > 0; i-- {

		val, _ := strconv.Atoi(string((id[index])))
		if i == 1 {
			last = val
		} else {
			sum += val * i
		}
		// fmt.Println(val)

		index++
	}

	fmt.Println(sum)
	check = sum % 11
	// fmt.Println(check)

	check = ((11 - check) % 10)
	fmt.Println(check)

	if check == last {
		fmt.Println("True")
		return true
	} else {
		fmt.Println("False")
		return false
	}

}

func main() {
	IsValid("1234567890121")
}
