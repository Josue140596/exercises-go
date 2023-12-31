package main

import (
	"exercise/slices/technicaltest"
)

func main() {

	//* Structs
	// structs.PrintShape(structs.Rectangle{Width: 10, Height: 20})
	// structs.PrintShape(structs.Circle{Radius: 10})
	// structs.PrintShape(structs.Triangle{A: 10, B: 20, C: 30})

	//* Pointers
	// pointers.BasicPointers()
	// pointers.PointerReceivers()

	//* Slices
	// slices.SlicesExamples()

	//* Contexts
	// contexts.ContextsExamplesWithTimeOut()
	// contexts.ContextWithValues()

	//* Routines and wait groups
	// routines.WithoutRoutines()
	// routines.WithRoutines()
	// routines.WithRoutinesWaitGroup()
	// routines.UseCaseRoutines()

	//* Idiom Ok
	// idiomok.WithoutOk()
	// idiomok.WithOk()
	// idiomok.UseCaseIdiomOk()

	//* Data races
	// dataraces.ExampleDataRace()
	// dataraces.WorkingWithMutex()
	// dataraces.WorkingWithAtomic()

	//============================================================
	//* Exercises - Slices - Use Case 01
	// slices.Prices()

	//* Technical Tests
	// Palindrome
	technicaltest.TTPalindrome()
}
