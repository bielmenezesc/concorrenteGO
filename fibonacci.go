package main

import (
	"fmt"
	"time"
)

func calma() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Calma Carlos!")
	}
}

func fibonacci(value int ) int {
	if value < 2 {
		return value
	} else {
		return fibonacci(value - 1) + fibonacci(value - 2)
	}
}

func main() {
	go calma()

	value := 0
	fmt.Scanln(&value)
	fmt.Println(fibonacci(value))
}
