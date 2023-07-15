package routine

import (
	"time"
	"fmt"
)

func processInput(input string) {
	// Simulate some time-consuming processing
	time.Sleep(2 * time.Second)

	// Display the processed result
	fmt.Println("Processed input:", input)
}

func handleUserInput() {
	fmt.Println("Enter some input:")

	// Start a goroutine to process the input
	go func() {
		var input string
		fmt.Scanln(&input)
		processInput(input)
	}()

	// Continue executing other tasks while waiting for input
	for {
		fmt.Println("Program is running...")
		time.Sleep(1 * time.Second)
	}
}