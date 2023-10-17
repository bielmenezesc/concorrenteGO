package main

import (
	"fmt"
	"math/rand"
	"time"
)

func request_stream() chan string {
	canal_strings := make(chan string)
	
	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyz"

	go func() {
		defer close(canal_strings)

		for i := 0; i < 100; i++ {

			randomStr := charset[rand.Intn(len(charset))]
			canal_strings <- string(randomStr)
		}
	}()

	return canal_strings
}

func ingest(in chan string, ch1 chan string, ch2 chan string) {
	for {
		select {
		case word := <-ch1:
			in <- word
		case word := <-ch2:
			in <- word
		}
	}
}

func main() {
	ch1 := request_stream()
	ch2 := request_stream()
	ch3 := make(chan string)
	
	go ingest(ch3, ch1, ch2)

	for word := range ch3 {
		fmt.Println(word)
	}
}
