/* Alta3 Research | Go Essentials
   Function literals - Anonymous     */

package main

import (
	"fmt"
	"sort"
)

func main() {

	// To create an anonymous struct
	// just instantiate the instance immediately after declaring the type
	people := []struct {
		Host string
		Mac  string
		Ram  int
		Cpu  int
	}{
		{"merlin", "2e549138a9e3", 1024, 1},
		{"prospero", "3c539188c9e3", 2048, 2},
		{"nestor", "1b166127a9e3", 2048, 1},
		{"odin", "4d545b88c9e3", 4096, 2},
	}
	people2 := people[:]
	// this technique makes sense if you're going to use a struct once and be done
	// marshaling or unmarshalling json, or prepping HTTP responses

	// In the following, if i is 0, j is 1 (i is the index before j)
	// The less function is a comparator function that reports whether the element at index i
	// should sort before the element at index j. If both less(i, j) and less(j, i) are false,
	// then the elements at index i and j are considered equal.
	//
	// An anonymous function is a function which doesn’t contain any name.
	// It is useful when you want to create an inline function.
	//                                                sort in ascending order
	sort.Slice(people, func(i, j int) bool { return people[i].Host < people[j].Host })
	fmt.Println("By Hostname:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Ram < people[j].Ram })
	fmt.Println("By RAM:", people)

	// assigning a nameless function to a variable is also possible
	algo := func(i, j int) bool { return people[i].Mac > people[j].Mac }
	sort.Slice(people, algo)
	fmt.Println("Reverse sort by MAC:", people)

	// Note: the sorting is done to the slices underlying array! Therefore, all other referencing slices will be sorted
	fmt.Println("people2: ", people2)

}
