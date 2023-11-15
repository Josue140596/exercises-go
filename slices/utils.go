package slices

import "slices"

// Remove by loop
// This function removes an element from a slice by index
// However, it is not efficient because it has to loop through all the elements
func removeByIndexLoop(slice []int, index int) []int {
	// Create a new slice with the same length as the original slice - 1
	result := make([]int, 0, len(slice)-1)
	for i, v := range slice {
		if i != index {
			result = append(result, v)
		}
	}
	return result
}

// Remove by copies
func removeNewSliceByCopies(slice []int, index int) []int {
	result := make([]int, len(slice)-1)
	// Copy the first part of the slice
	// For example, if index = 2, then copy the first 2 elements
	copy(result, slice[:index]) // [10 20]
	// Copy the second part of the slice
	// For example, if index = 2, then copy the last 6 elements
	copy(result[index:], slice[index+1:]) // [40 50 60 70 80 90]
	return result
}

// Remove by copy
func removeNewSliceByCopy(slice []int, index int) []int {
	// Copy the first part of the slice
	// For example, if index = 2, then copy the first 2 elements
	copy(slice[index:], slice[index+1:]) // [10 20 40 50 60 70 80 90]
	return slice[:len(slice)-1]          // [10 20 40 50 60 70 80 90]
}

// Remove by append
func removeByAppend(slice []int, index int) []int {
	// Append the first part of the slice
	// For example, if index = 2, then append the first 2 elements
	// [10 20]
	// Append the second part of the slice
	// For example, if index = 2, then append the last 6 elements
	// [40 50 60 70 80 90]
	//... is the ellipsis operator which expands the slice into individual elements
	return append(slice[:index], slice[index+1:]...)
}

// Remove by slices package
func removeBySlices(slice []int, index int) []int {
	return slices.Delete(slice, index, index+1)
}
