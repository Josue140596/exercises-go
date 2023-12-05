package benchmarks

import (
	"exercise/slices/slices"
	"fmt"
)

func BenchmarksExamples() {
	//Slice is a reference type
	//A slice is a reference to an underlying array
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("x: ", x) // [10 20 30 40 50 60 70 80 90]

	//Slicing
	//Initial position index is included 0 , 1 , 2 , 3 , 4 , 5 , 6 , 7 , 8
	//Final position index is excluded 1 , 2 , 3 , 4 , 5 , 6 , 7 , 8 , 9
	fmt.Println("x[1:5]: ", x[1:3]) // [20 30 40]
	//You can omit the initial position index
	fmt.Println("x[:5]: ", x[:5]) // [10 20 30 40 50]
	// If you omit the final position index, it will slice until the end of the slice
	fmt.Println("x[5:]: ", x[5:]) // [60 70 80 90]

	// 1. Remove by loop (inefficient)
	fmt.Println("Remove by loop", slices.RemoveByIndexLoop(x, 2)) // [10 20 40 50 60 70 80 90]
	// 2. Remove by copy (efficient but not readable)
	fmt.Println("Remove by copies", slices.RemoveNewSliceByCopies(x, 2)) // [10 20 40 50 60 70 80 90]
	// 3. Remove by copy (efficient and readable)
	fmt.Println("Remove by copy", slices.RemoveNewSliceByCopy(x, 2)) // [10 20 40 50 60 70 80 90]
	// 4. Remove by append (efficient and readable)
	fmt.Println("Remove by append", slices.RemoveByAppend(x, 2)) // [10 20 40 50 60 70 80 90]
	// 5. Remove by slices package (efficient and readable)
	fmt.Println("Remove by slices package", slices.RemoveBySlices(x, 2)) // [10 20 40 50 60 70 80 90]
}
