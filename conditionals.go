package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Basic if conditionals using "Simple Statements"
	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("First course is doing better than second course.")
	} else if firstRank > secondRank {
		fmt.Println("Second course doing better first course.")
	} else {
		fmt.Println("Both courses are tied.")
	}

	// Basic switch conditional
	// NB: 'fallthrough' keyword will execute a case after the matched condition
	switch "docker" {
	case "linux":
		fmt.Println("Here are some recommended Linux courses:")
	case "docker":
		fmt.Println("Here are some recommended Docker courses:")
	default:
		fmt.Println("Match not found.")
	}

	// Breaking switch conditional
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Even number - ", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd number -", tmpNum)
	}

	// Error checking
	// NB: Make function return error - func f1(s1 string) (string, error)
	_, err := os.Open("/example.md")
	if err != nil {
		fmt.Println("An error was returned: ", err)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
