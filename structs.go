package main

import "fmt"

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta) - Same as above but gives a ptr

	// Can also be intialised without fields - courseMeta{"Nigel", "Int.." "5"}
	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println("Docker Deep Dive author is:", DockerDeepDive.Author)
	DockerDeepDive.Rating = 3
	fmt.Println(DockerDeepDive)
}
