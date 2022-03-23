/* Alta3 Research | RZFeeser
   Flags - Using the init() function   */

package main

import (
	"flag"
	"fmt"
)

var (
	env1  *string
	port1 *int
)

// the 'init()' function is often used to initialize state variables
func init() {
	env1 = flag.String("env1", "development", "current environment")
	port1 = flag.Int("port1", 3000, "port1 number")
}

func main() {

	flag.Parse()

	fmt.Println("env1:", *env1)
	fmt.Println("port1:", *port1)
}
