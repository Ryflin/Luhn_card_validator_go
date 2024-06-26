package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)
var url = "http://localhost:8080/"
type validateRequest struct {
	Name string `json:"Password"`
	Password string `json:"client_id"`
	Number int `json:"CardNumber"`
}
type validateResponse struct {
	Valid bool `json:"valid"` 
	Limit uint64	`json:"limit"`
}


func main() {
	var number = read_digit()
	for luhn_digit(number) {
		go validate_number(number)
	}
}


// validates an account wtih the server and checks to see if the balance should be declined
func validate_with_server(username, password, number string) (valid bool, limit uint64, err error) {
	valid = false
	limit = 0
	hasher := sha256.New()
	var temp validateRequest
	temp.Name = username
	temp.Number, err = strconv.Atoi(number)
	if err != nil {
		log.Println("Error while converting number: ", err.Error())
		return valid, limit, fmt.Errorf("error: can try again with different values")
	}
	hasher.Write([]byte(password))
	temp.Password = hex.EncodeToString(hasher.Sum(nil))

	// make into json
	jsonData, err := json.Marshal(temp)
	if err != nil {
		log.Println("Error while Unmarshalling ", err.Error())
		err = fmt.Errorf("error, can try different values")
		return valid, limit, err
	}
	// read bytes into post
	resp, err := http.Post(url + "validateUser", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Println("Error making POST Request")
		return valid, limit, fmt.Errorf("error: unable to connect")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Recieved %d as an error code", resp.StatusCode)
		err = fmt.Errorf("error: try again later")
		return valid, limit, err
	}
	var response_struct validateResponse
	var response_data []byte
	response_data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error with reading Response: ", err.Error())
	}
	err = json.Unmarshal(response_data, &response_struct)
	if err != nil {
		log.Println("Error unmarshalling response", err)
		err =  fmt.Errorf("error on backend can try again")
		return valid, limit, err
	}
	return response_struct.Valid, response_struct.Limit, nil
}

func check_card_Provider() {

}

func validate_number(number string) {
	var valid = false

	
	if valid {
		println("Number ", number, " is valid")
	} else {
		println("Number ", number, " isn't valid")
	}
}



func read_digit() (number string) {
	var err = fmt.Errorf("put in a number: ")
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
	if len(digits) < 2 {
		return false;
	}
	digits = strings.ReplaceAll(digits, " ", "")
	var digit = 0
	var even = true
	for i := len(digits) - 2; i >= 0; i-- {
		temp_digit, err := strconv.Atoi(string(digits[i]))
		print(temp_digit) // debugging remove later
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
	println()
	last, err := strconv.Atoi(string(digits[len(digits) - 1]))
	if err != nil {
		println("error: ", err.Error())
		return false
	}
	println(last, " + ", digit)
	temp_dig := 10 - digit % 10

	if temp_dig != last {
		// println("error: ", err.Error())
		return false
	}
	return true
}