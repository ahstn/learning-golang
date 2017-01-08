package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic for loops
	for i := 10; i >= 0; i-- {
		if i == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	// For range loops
	courseInProgress := []string{"Docker Deep Dive", "Docker Clustering",
		"Docker and Kubernetes"}

	for _, i := range courseInProgress {
		fmt.Println(i)
	}

	// Using continue in loops
	// Will only print odd numbers
	for i := 10; i >= 0; i-- {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
