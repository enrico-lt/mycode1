package main

import "fmt"

type Virtmach struct {
	ip       string
	hostname string
	diskGb   int
	ram      int
}

func (v *Virtmach) getIp() string {
	return v.ip
}

func (v *Virtmach) setRam(ram int) {
	v.ram = ram
}

func (v *Virtmach) toString() string {
	return fmt.Sprintf("%v", v)
}

func main() {

	v1 := Virtmach{
		ip:       "127.0.0.1",
		hostname: "localhost",
		diskGb:   100,
		ram:      16,
	}
	fmt.Println(v1.toString())

	v1.setRam(32)

	fmt.Println(v1.toString())

}
