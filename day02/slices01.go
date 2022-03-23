/* RZFeeser | Alta3 Research
   Slices relationship to arrays     */

package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5} // this array will back the slice
	slice := array[:]              // [:] means "from start to the end of"
	slice2 := array[:]             // [:] means "from start to the end of"
	slice3 := make([]int, len(array))
	copy(slice3, slice)

	//Modifying the slice (which is a pointer to the array)
	slice[1] = 7
	fmt.Println("Modifying Slice")
	fmt.Println("Array =", array)
	fmt.Println("Slice =", slice)
	fmt.Println("Slice2 =", slice2) // also, slice2 is updated!!!
	fmt.Println("Slice3 =", slice3) // slice 3 not updated because deep(?) copy of slice

	//Modifying the array. Would reflect back in slice too
	array[1] = 2
	fmt.Println("Modifying Underlying Array")
	fmt.Println("Array =", array)
	fmt.Println("Slice =", slice)
	fmt.Println("Slice2 =", slice2) // also, slice2 is updated!!!
	fmt.Println("Slice3 =", slice3) // also, slice3 is updated!!!
}
