package client

// damm gurl

import (
	"fmt"
	"strconv"
)

var matrix = [10][10]int{
	{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
	{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
	{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
	{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
	{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
	{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
	{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
	{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
	{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
	{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
}

// performs the damm algorithm and appends the check digit to the string
func dammEncode(number string) (check string) {
	var interm = 0
	for i := 0; i < len(number); i++ {
		temp, _ := strconv.Atoi(string(number[i]))
		interm = matrix[interm][temp]
	}
	number += fmt.Sprint(interm)
	return number
}

func dammDecode(number string) (correct bool) {
	temp := dammEncode(number)
	temp = string(temp[len(temp) - 1])
	return temp == "0" 
}