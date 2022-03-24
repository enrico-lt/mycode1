/* Alta3 Research | RZFeeser
   Consuming RESTful APIs */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Capsule []struct {
	CapsuleSerial      string    `json:"capsule_serial"`
	CapsuleID          string    `json:"capsule_id"`
	Status             string    `json:"status"`
	OriginalLaunch     time.Time `json:"original_launch"`
	OriginalLaunchUnix int       `json:"original_launch_unix"`
	Missions           []struct {
		Name   string `json:"name"`
		Flight int    `json:"flight"`
	} `json:"missions"`
	Landings   int    `json:"landings"`
	Type       string `json:"type"`
	Details    string `json:"details"`
	ReuseCount int    `json:"reuse_count"`
}

func main() {

	// define our URL (API) as a string
	url := "https://api.spacexdata.com/v3/capsules"
	// the HTTP method we wish to send (HTTP GET)
	method := "GET"

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// build a new HTTP request
	req, err := http.NewRequest(method, url, nil)

	// error checking
	if err != nil {
		fmt.Println(err)
		return
	}

	// if the request requires basic AUTH
	// (a protected resource), use the following:
	//                  //username  //password
	// req.SetBasicAuth("userRZF", "passQwerty!")

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close() // even if we have a problem, we still want this to happen

	//
	body, err := ioutil.ReadAll(res.Body)

	// error checking
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	//fmt.Println(string(body))

	// we initialize our Capsules array
	var capsules Capsule

	err = json.Unmarshal(body, &capsules)
	if err != nil {
		fmt.Printf("Error unmarshallign!! %v", err)
	}

	//for i := 0; i < len(capsules); i++ {
	//	fmt.Println("capsule: ", capsules[i])
	//}

	//for _, c := range capsules {
	//	fmt.Printf("capsule: %v\n", c)
	//}

	// Other method for unmarshalling:

	// Use json.Decode for reading streams of JSON data
	// -> use the response directly and not converting it into []byte with ioutil.ReadAll(res.Body)
	if err := json.NewDecoder(res.Body).Decode(&capsules); err != nil {
		log.Println(err)
	}

	// the _ is an int
	// launchData is the data
	for launchNo, launchData := range capsules {
		fmt.Println("Capsule Record     =\n", launchData)
		fmt.Println("Record Number      =", launchNo)
		fmt.Println("Capsule Serial     =", launchData.CapsuleSerial)
	}

}
