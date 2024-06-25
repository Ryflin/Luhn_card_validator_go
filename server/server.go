package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type card struct {
	Number int `json:"Number"`
	CardLimit uint64 `json:"Limit"`
	CurrentlySpent uint64 `json:"Spent"`
	CreditScore int `json:"Score"`
	User string `json:"User"`
	Password string `json:"Client_id"`
	Provider string `json:"Provider"`
}

type validateRequest struct {
	Name string `json:"Password"`
	Password string `json:"client_id"`
	Number int `json:"CardNumber"`
}

type validationResponse struct {
	Valid bool `json:"valid"`
	Limit uint64 `json:"limit"`
}



func getUserHashFile(userName string) (fileName string, exists bool) {
	return "db.json", true
}


func getUserCard(filename string, userHash string) (userCard card, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Debug("Cannot read user file: "+err.Error())
		return card{}, err
	}
	var allusers map[string]card
	err = json.Unmarshal(data, &allusers)
	if err != nil {
		return card{}, err
	}
	userCard, ok := allusers[userHash]
	if !ok {
		err = fmt.Errorf("error: cannot find user card in users")
		return card{}, err
	}
	return userCard, nil
}

func validateNumber(c *fiber.Ctx) (err error) {
	var resp validationResponse
	resp.Limit = 0
	resp.Valid = false
	log.Debug("testing")
	hasher := sha256.New()
	var recievedData validateRequest
	err = json.Unmarshal(c.Body(), &recievedData)
	if err != nil {
		println("error while unmarshalling")
		log.Debug("Error while unmarshalling from c.Body", err.Error())
		return err
	}
	recievedData.Name = strings.ToValidUTF8(recievedData.Name, "")
	hasher.Write([]byte(recievedData.Name))
	log.Debug(recievedData.Name)
	userHash := hex.EncodeToString(hasher.Sum(nil))
	log.Debug(userHash)
	var userCard card
	if filename, ok := getUserHashFile(userHash); ok {
		userCard, err = getUserCard(filename, userHash)

	} else {
		log.Debug("incorrect credentials")
	}
	if err != nil {
		log.Debug(err)
		err = c.JSON(resp)
		return err
	}
	log.Debug(recievedData.Name + " " + recievedData.Name + " " + recievedData.Password + " extra " + userCard.Password + " " + userCard.Password)
	if recievedData.Name != userCard.User || recievedData.Password != userCard.Password || recievedData.Number != userCard.Number {
		log.Debug("Incorrect credentials\n")
		err = c.JSON(resp)
		return err
	} 



	// testing response 
	resp.Limit = userCard.CardLimit - userCard.CurrentlySpent
	log.Debug("limit ", userCard.CardLimit, "  ", userCard.CurrentlySpent)
	resp.Valid = true
	err = c.JSON(resp)
	if err != nil {
		log.Debug(err.Error())
	}
	return
}


func main() {
	f := fiber.New()
	f.Post("/validateUser", validateNumber)
	log.Fatal(f.Listen(":8080"))
}