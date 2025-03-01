//Without Using Goroutine Channel:
//------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random generator

	totalPeopleCrossed := 0
	totalTime := 0

	fmt.Println("🏞️ Park Rounds Simulation 🏃‍♂️")

	for round := 1; round <= 5; round++ {
		peopleCrossed := rand.Intn(10) + 1 // Random people between 1-10
		timeTaken := rand.Intn(5) + 3      // Random time between 3-7 seconds

		totalPeopleCrossed += peopleCrossed
		totalTime += timeTaken

		fmt.Printf("Round %d: Crossed %d people, Time taken: %d seconds\n", round, peopleCrossed, timeTaken)
	}

	fmt.Println("\n🏁 Simulation Complete! 🏁")
	fmt.Printf("Total People Crossed: %d\n", totalPeopleCrossed)
	fmt.Printf("Total Time Taken: %d seconds\n", totalTime)
}
