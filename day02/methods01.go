/* Alta3 Research | RZFeeser
   Methods - defining         */

// Go program to illustrate the
// method with struct type receiver
package main

import "fmt"

// Author structure
type author struct {
	name   string
	branch string
	books  int
	salary int
}

// Method with a receiver
// of author type
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.books)
	fmt.Println("Salary: ", a.salary)
}

func show1(a author) {
	fmt.Println("function: Author's Name before: ", a.name)
	a.name = "new name"
	fmt.Println("function: Author's Name before: ", a.name)
}

func (a *author) bookup() {
	fmt.Println("method: Author's books before: ", a.books)
	a.books++
	fmt.Println("method: Author's books after: ", a.books)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name:   "RZFeeser",
		branch: "Pennsylvania",
		books:  14,
		salary: 34000,
	}

	// Calling the method
	res.show()
	fmt.Println("----------------")
	show1(res) // not changing the struct
	fmt.Println("after func\n", res)
	fmt.Println("----------------")
	res.bookup() // with the derefence * it modifies the caller
	fmt.Println("after method\n", res)

}
