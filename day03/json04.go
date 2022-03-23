package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	jsonFile, err := os.Open("C:\\data\\repo\\mycode1\\day03\\users.json")
	if err != nil {
		log.Fatalf("failed open: %s", err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	//var users Users
	//users := make(map[string]any)
	var users map[string]any

	json.Unmarshal([]byte(byteValue), &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	//for i := 0; i < len(users.Users); i++ {
	//	fmt.Println("UserCsv Type: " + users.Users[i].Type)
	//	fmt.Println("UserCsv Age: " + strconv.Itoa(users.Users[i].Age))
	//	fmt.Println("UserCsv Name: " + users.Users[i].Name)
	//	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	//}

	for key := range users {
		//fmt.Println("Key:", key, "=>", "Element:", element)
		fmt.Println(users[key])
	}
}
