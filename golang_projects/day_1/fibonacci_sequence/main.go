package main

import (
	fib "fibonacci_sequence/fibonacci"
	"fmt"
)

func main() {
	seq := fib.SequenceInitializer()
	fmt.Println("Which element of the Fibonacci Sequence would you like to be returned? ")
	var inputPosition int
	_, err := fmt.Scanf("%d", &inputPosition)
	if err != nil {
		panic(err.Error())
	}
	element := seq.FibonacciElement(inputPosition)
	fmt.Printf("The element you're looking for is %d\n", element)
}
