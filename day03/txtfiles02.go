/* RZFeeser | Alta3 Research
   Reading text files with bufio package - writing contents into a slice
   See documentation @ https://pkg.go.dev/bufio  */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// os.Open() is platform-independent function
	file, err := os.Open("C:\\data\\go\\schulung\\txtfiles01.txt") // for read access

	// error checking
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	// ensure the file is closed
	defer file.Close()

	// this snippet will reproduce the file on the screen
	scanner := bufio.NewScanner(file)

	// we could also write the data into a slice
	var txtlines []string // create a slice of strings
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text()) // each time through the loop build the slice
	}

	// now we have a slice that contains the contents of our file
	// we can loop across that file and reproduce it on the screen
	// via the following snippet
	for _, eachline := range txtlines { // the _ would be the "line number" (starting at 0)
		fmt.Println(eachline)
	}

	// write to an other file
	fileWrite, err := os.OpenFile("C:\\data\\go\\schulung\\txtfiles03.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// error checking
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	dataWriter := bufio.NewWriter(fileWrite)

	for _, eachLine := range txtlines {
		_, err := dataWriter.WriteString(eachLine + "\n")
		if err != nil {
			return
		}
	}

	dataWriter.Flush()
	fileWrite.Close()
}
