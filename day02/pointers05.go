/* RZFeeser | Alta3 Research
   With Pointers - Why we need pointers  */

package main

import (
	"fmt"
)

type User1 struct {
	Name string
	Pet  []string
}

// a method for type User1
func (p2 *User1) newPet() {
	fmt.Println(*p2, "underlying value of p2 before")
	p2.Pet = append(p2.Pet, "Fido")
	fmt.Println(*p2, "underlying value of p2 after")
}

func main() {
	u := User1{Name: "Anna", Pet: []string{"Bailey"}} // this time we'll generate a pointer!

	fmt.Println(u, "u before")

	p := &u // Let's make a pointer to u!

	p.newPet()
	//u.newPet() // also works
	fmt.Println(u, "u after") // Does Anna have a new pet now? Yes!
}
