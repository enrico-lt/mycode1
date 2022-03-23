/* Alta3 Research | RZFeeser
   Reading input with Scanf   */

package main

import "fmt"

func main() {

	var name string // define a string variable (for tracking user name)

	fmt.Print("Enter your name: ") // "fmt.Print" is not followed by a /n
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return
	} // %s - string, and a pointer to "name" (otherwise we'll make a copy of name and change the copy)
	fmt.Println("Hello", name) // Print "Hello" followed by the value of name
}
