package main

import "fmt"

func main() {
	/* Maps like slices are reference types (passed to funcs by ref)
	   Therefore they're unsafe for concurrency
	   map[keyType]valueType, size (keyType must be a comparable)
	   Size is optional but can improve performance */

	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}

	fmt.Printf("League titles: %v\nRecent head to head: %v\n", leagueTitles,
		recentHead2Head)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	for key, value := range testMap {
		fmt.Printf("\nKey is: %v, Value is: %v", key, value)
	}

	testMap["A"] = 100
	testMap["F"] = 56
	delete(testMap, "E")
	fmt.Println("\n", testMap)
}
