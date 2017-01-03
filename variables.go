package main

import (
	"fmt"
	"os"
	"reflect"
)

var (
	name   = os.Getenv("USER")
	course = "Docker Deep Dive" // NB: Won't compile when declared inside a function as it's not used.
	module = 3.2
)

func main() {
	// Type checking and infering
	fmt.Println("Name is:", name, "and is of type:", reflect.TypeOf(name))
	fmt.Println("Module is:", module, "and is of type:", reflect.TypeOf(module))

	// Type conversion
	a := 10.0
	b := 3
	c := int(a) + b

	fmt.Println("\nA is type:", reflect.TypeOf(a), "and B is of type:", reflect.TypeOf(b))
	fmt.Println("Module is:", c, "and is of type:", reflect.TypeOf(c))

	// Pointers - NB: '&' references a pointer and '*' de-references it
	ptr := &module
	fmt.Println("\nMemory address of $module is:", ptr, "and the value is:", *ptr)
	fmt.Println("\nHi", name, ", you're currently watching:", course)

	changeCourse(&course)

	fmt.Println("You are now watching course:", course)

	// Constants - NB: Can't be defined using ':=' shorthand
	const speedOfLight = 186000

	// Environment Variables
	// fmt.Println(os.Environ()) - Prints a messy list of env vars

	//for _, env := range os.Environ() { - Prints each env var on a seperate line
	//fmt.Println(env)
	//}
}

// NB: Removing '*' and '&' will only access a value copy of course, not the global definition
func changeCourse(course *string) string {
	*course = "First Look: Native Docker Clustering"
	fmt.Println("Trying to change your course to:", *course)

	return *course
}
