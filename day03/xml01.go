/* RZFeeser | Alta3 Research
Exploring arrays              */
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Subject string `xml:"subject"`
	Body    string `xml:"body"`
}

func main() {
	data, _ := ioutil.ReadFile("C:\\data\\repo\\mycode1\\day03\\avengers.xml") // open the file "avengers.xml"

	note := &Notes{}

	_ = xml.Unmarshal([]byte(data), &note)

	fmt.Println(note.To)      // display value of "to"
	fmt.Println(note.From)    // display value from "from"
	fmt.Println(note.Subject) // display value from "subject"
	fmt.Println(note.Body)    // display value from "body"

	note1 := Notes{
		To:      "Kevin McCallister",
		From:    "Mom",
		Subject: "Be home soon!",
		Body:    "Getting the first flight home! Take care of the house until we return.",
	}

	file, _ := xml.MarshalIndent(note1, "", " ")

	_ = ioutil.WriteFile("C:\\data\\repo\\mycode1\\day03\\HomeAlone.xml", file, 0644)
}
