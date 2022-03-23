package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("C:\\data\\go\\schulung\\txtfiles01.txt")
	if err != nil {
		log.Fatalf("failed open: %s", err)
	}
	// ensure the file is closed
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
