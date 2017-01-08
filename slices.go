package main

import "fmt"

func main() {
	/* Slices are chunks of an array.
	   Each index in a slice is a pointer reference to the array index
	   Due to this, it can't be bigger than the referenced array */

	/* Declaring a slice - Go creates an array in the background to store data
	   make(<type>, <length>, <capacity>)
	   Length being slice size
	   Capacity being the size of the underlying array (optional param) */
	slice := make([]string, 5, 10)
	fmt.Printf("Length is: %d \nCapacity is: %d", len(slice), cap(slice))

	// Declaring and initialising a slice in 1 line
	// Now the slice is set to length 3 and the capacity is also 3
	myCourses := []string{"Docker", "Puppet", "Chef", "Python", "Ruby"}
	fmt.Println("\n\n2nd Course is:", myCourses[1])
	fmt.Println("All Courses are:", myCourses)

	/* Defining subset slices
	   [1:4] - 2nd item to 4th item
	   [:4]  - 1st item to 4th item
	   [2:]  - 3rd item to last item */
	subCourses := myCourses[:3]
	fmt.Println("The first 3 courses are:", subCourses)

	// Adding beyond the length of a slice
	// Go automatically handles this and increases the array size by 2^x
	intSlice := make([]int, 1, 4)
	fmt.Println("\n", intSlice)
	fmt.Printf("Length is: %d, Capacity is: %d\n", len(intSlice), cap(intSlice))

	for i := 1; i < 17; i++ {
		intSlice = append(intSlice, i)
		fmt.Printf("\nCapacity is: %d", cap(intSlice))
	}
	fmt.Println("\n", intSlice)

	// for range returns 2 values - index and data (as seen in loops.go:23)

	// Slices can be appended to each other by: append(slice1, slice2...)
}
