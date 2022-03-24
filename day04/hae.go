package main

import "fmt"

var myInt int

type myObject struct {
	name string
}

func myFuncVal(i int) {
	fmt.Println("In myFuncVal changing input")
	i = 9
}

func myFuncRef(i *int) {
	fmt.Println("In myFuncRef changing input")
	*i = 9
}

func (obj *myObject) myObjectMethodReturnRef(name string) *myObject {
	obj.name = name
	return obj
}

func (obj myObject) myObjectMethodReturnVal(name string) myObject {
	obj.name = name
	return obj
}

func (objRef *myObject) myObjectMethodRefReturnNil(name string) {
	objRef.name = name
}

func (objVal myObject) myObjectMethodValReturnNil(name string) {
	objVal.name = name
}

func main() {

	var myIntPointer *int = &myInt

	myInt = 1
	fmt.Println(myInt)
	fmt.Println(myIntPointer)
	fmt.Println("adress with &: ", myIntPointer)
	fmt.Println("Dereferenced with *: ", *myIntPointer)

	// modify the value behind the pointer
	*myIntPointer = 2
	fmt.Println("Update of *myIntPointer: ", *myIntPointer)
	fmt.Println("Also underlying int changed: ", myInt)

	myFuncVal(myInt)
	fmt.Println("After myFuncVal: ", myInt)

	myFuncRef(&myInt)
	fmt.Println("After myFuncRef: ", myInt)

	o1 := myObject{name: "aName"}
	oResult := o1.myObjectMethodReturnRef("newName")
	fmt.Println("After myObjectMethodReturnRef: ", *oResult, "with input obj: ", o1)

	o2 := myObject{name: "bName"}
	oResultVal := o2.myObjectMethodReturnVal("newName")
	fmt.Println("After myObjectMethodReturnVal: ", oResultVal, "with input obj: ", o2)

	o3 := myObject{name: "cName"}
	o3.myObjectMethodRefReturnNil("newName")
	fmt.Println("After myObjectMethodRefReturnNil: ", o3, "with input obj: ", o3)

	o4 := myObject{name: "dName"}
	o4.myObjectMethodValReturnNil("newName")
	fmt.Println("After myObjectMethodValReturnNil: ", o4, "with input obj: ", o4)

	var oRef *myObject
	oRef = &o3
	//oRef.myObjectMethodValReturnNil("gaga") // -> nothing is changed on o3 and oRef
	oRef.myObjectMethodRefReturnNil("gaga")
	fmt.Println("oRef: ", *oRef)
	fmt.Println("now o3: ", o3)

}
