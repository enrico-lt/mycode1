package main

import "fmt"

type Virtmach struct {
	ip       string
	hostname string
	diskGb   int
	ram      int
}

func (v *Virtmach) GetIp() string {
	return v.ip
}

func (v *Virtmach) SetRam(ram int) {
	v.ram = ram
}

func (v *Virtmach) ToString() string {
	return fmt.Sprintf("%v", v)
}
