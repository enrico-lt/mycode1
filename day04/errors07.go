/* Alta3 Research | RZFeeser
   Handling errors within our models
      - Create a type
      - Wrap the initial value
      - Wrap the behavior in the type
      - Wrap the conditionals into the type */

package main

import (
	"errors"
	"fmt"
)

// create a new type
type Superdivider struct {
	answer int
	err    error
}

// wrap the behavior
func (s *Superdivider) divide(x int) {
	// wrap the conditionals (create our guard)
	if s.err != nil {
		return
	} else if x == 0 {
		s.err = errors.New("math: Sorry, you cannot div by 0.") // raise an error
		return
	}
	// protected by our guard, we can now perform the operation
	s.answer = s.answer / x
}

// this function controls the "level of abstraction" we want
// to offer our user
func divide(a, b, c int) {
	// wrap the initial value
	z := Superdivider{a, nil}
	z.divide(b)
	z.divide(c)

	// determine result
	if z.err != nil {
		//panic(z.err)
		fmt.Println(z.err)
		return
	}
	fmt.Println(z.answer)
}

func main() {
	// run divide
	divide(100, 2, 2) // should return 25

	// run divide by zero
	divide(100, 0, 20) // this causes an error
	// but our handler has us at the same spot

	// run divide by zero (again)
	divide(100, 5, 0) // we are still handling our errors, just in a highly
	// abstracted way
}
