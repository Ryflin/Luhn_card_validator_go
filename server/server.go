package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

type card struct {
	Number int `json:"number"`
	CardLimit uint64 `json:"Limit"`
	CurrentlySpend uint64 `json:"Spent"`
	CreditScore int `json:"Score"`
	User string `json:"user_name"`
	Password string `json:"Client_id"`
	Provider string `json:"provider"`


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

type account struct {
	Holder map[string]card
}


// func UpdateAccount(args...account) (account, error){
	
// 	if len(args) > 1 {
// 		err := fmt.Errorf("Only one account can be made at a time")
// 	}
// }



// func createNewCard(args...string) {

// }



func getUserHashFile(userName string) () {
	return 
}


func validateNumber(c *gin.Context) {
	var recievedData validateRequest
	jsonData, err := c.GetRawData()
	if err != nil {
		log.Println("Error getting raw data: ", err)
	}
	err = json.Unmarshal(jsonData, &recievedData)
	
	// testing response 
	var resp validationResponse
	resp.Limit = 200
	resp.Valid = true
	c.JSON(200, resp)
}

// func rateLimit()

func main() {
	g := gin.New()
	g.POST("validateNumber", validateNumber)
	// g.Use()
	g.Run()
}