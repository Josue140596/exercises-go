package idiomok

import "fmt"

func WithoutOk() {
	var myVar interface{}
	myVar = 1.5

	//Now we need to work with the value of myVar as an int
	//Because we don't know if myVar is an int, we need to check it
	var sum int
	// sum = myVar + 1 //This will fail because myVar is an interface
	a := myVar.(int) + sum

	fmt.Printf("a: %v\n", a)

}
