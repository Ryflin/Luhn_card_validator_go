package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var number = read_digit()
	for luhn_digit(number) {

	}
}

func read_digit() (number string) {
	var err = fmt.Errorf("Put in a number: ")
	for err != nil {
		println(err.Error() + " ")
		_, err = fmt.Scanln(number)
	}
	return number
}

// takes in a digit as a string and performs luhn validation on it
//
// `digits` the "credit card" string with the last number being the payload
//
// Information about luhn validation found here https://en.wikipedia.org/wiki/Luhn_algorithm
func luhn_digit(digits string) (valid bool) {
	digits = strings.ReplaceAll(digits, " ", "")
	var digit = 0
	var even = true
	for i := len(digits) - 2; i > 0; i-- {
		temp_digit, err := strconv.Atoi(string(digits[i]))
		println(temp_digit) // debugging remove later
		if err != nil {
			println()
			return false

		}
		if even {
			temp_digit *= 2
			if temp_digit > 10 {
				temp_digit = temp_digit - 9
			}
		}
		even = !even
		digit += temp_digit
	}
	last, err := strconv.Atoi(string(digits[len(digits) - 1]))
	if err != nil || 10 % digit != last {
		return false
	}
	return true
}