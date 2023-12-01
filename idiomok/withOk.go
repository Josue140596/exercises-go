package idiomok

import "fmt"

func WithOk() {
	var myVar interface{}
	myVar = 1

	//Now we need to work with the value of myVar as an int
	//Because we don't know if myVar is an int, we need to check it
	var sum int
	// sum = myVar + 1 //This will fail because myVar is an interface

	//Ok is a boolean that indicates if the assertion was successful
	//If ok is false, then a will be the zero value of the type
	//If ok is true, then a will be the value of the type
	val, ok := myVar.(int)

	if ok {
		sum = val + 1
		fmt.Printf("sum: %v\n", sum)
	} else {
		fmt.Printf("myVar is not an int\n")
	}

}
