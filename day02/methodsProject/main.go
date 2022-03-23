package main

import "fmt"

func main() {

	v1 := Virtmach{
		ip:       "127.0.0.1",
		hostname: "localhost",
		diskGb:   100,
		ram:      16,
	}
	fmt.Println(v1.ToString())

	v1.SetRam(32)

	fmt.Println(v1.ToString())

}
