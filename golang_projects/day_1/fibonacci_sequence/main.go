package main

import (
	"fibonacci_sequence/fibonacci"
	"fibonacci_sequence/handlers"
	"fmt"
)

func main() {
	seq := fibonacci.SequenceInitializer()
	for {
		partialSequence, err := handlers.AppInteraction(&seq)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("The element you're looking for is %d\n", partialSequence)
	}

}
