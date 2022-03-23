/* Alta3 Research | RZFeeser
   Flags - Strings and Integers  */

package main

import (
	"flag"
	"fmt"
)

var (
	env  *string
	port *int
)

func main() {

	env := flag.String("env1", "development", "current environment")
	port := flag.Int("port1", 3000, "port1 number")

	flag.Parse()

	fmt.Println("The value of env1:", *env)          // display the dereferenced values
	fmt.Println("The port1 has been set to:", *port) // display the dereferenced values
}
