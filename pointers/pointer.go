package pointers

import "fmt"

func BasicPointers() {
	var myVar string = "Original Value"

	//What is a pointer?
	//A pointer is a variable which stores the memory address of another variable
	var ptr *string = &myVar
	fmt.Printf("Value of myVar: %v\n", myVar)
	fmt.Printf("Dir of myVar: %v\n", &myVar)
	fmt.Printf("Value of ptr: %v\n", ptr)
	fmt.Printf("Dir of ptr: %v\n", &ptr)
	//Another important thing to note is that
	//You can get the value of a pointer by using the * operator
	fmt.Printf("Value of ptr through the pointer (*ptr) : %v\n", *ptr)
	//Modify the value of myVar through the pointer
	//Get the value of myVar through the pointer
	*ptr = "Modified Value"
	fmt.Printf("Value modified of ptr through the pointer (*ptr) : %v\n", *ptr)
	//Now the value of myVar is modified too.
	fmt.Printf("Value MODIFIED of >>myVar<<: %v\n", myVar)
}

func PointerReceivers() {
	//When I should use a pointer receiver?
	//First we have two scenarios:
	//*1. You want to modify the value of the receiver by that method.
	player1 := player{points: 0}

	fmt.Printf("Player points: %v\n", player1.points)
	player1.addPoints(100)
	// The value of player1.points should be 100, isn't it?
	fmt.Printf("Player AFTER points added with addPoints method: %v\n", player1.points)
	// But it's not, because the addPoints method is using a copy of the player1 struct.
	// And they're using different memory addresses.
	// So, how can we modify the value of the receiver by that method?
	// We can use a pointer receiver.
	// A pointer receiver uses a pointer to the original struct.
	// So, when we modify the value of the pointer receiver, we're modifying the original struct.
	// Let's see how it works.
	player1.addPointsPointerReceiver(100)
	// The value of player1.points going to be 100
	fmt.Printf("Player AFTER points added with addPointsPointerReceiver method: %v\n", player1.points)

	//*2. If you receiver is a large struct, it will be more efficient to use a pointer receiver.
	//For instance, if you have a struct with 100 fields or slice with many records, it will be more efficient to use a pointer receiver.
	//Because you're not copying the struct or the slice.
}

// 1. You want to modify the value of the receiver by that method.
type player struct {
	points  int
	records []string
}

// Remember that this method is using a copy of the player struct.
func (p player) addPoints(points int) {
	p.points += points
}

// Now, this method is using a pointer to the original struct.
func (p *player) addPointsPointerReceiver(points int) {
	p.points += points
}
