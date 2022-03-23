/* RZFeeser | Alta3 Research
   switch - multiple expressions */

package main

import (
	"fmt"
	"time"
)

func main() {

	switch time.Now().Weekday() {

	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Working for the weekend.")
	case time.Saturday, time.Sunday:
		fmt.Println("Party time!")
	}
}
