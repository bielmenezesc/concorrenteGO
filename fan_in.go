package main

import (
	"fmt"
	"time"
)

func request_stream() chan string {
	canal_strings := make(chan string)

	go func() {
		defer close(canal_strings)

		for i := 0; i < 100; i++ {

			canal_strings <- "abacate"
		}
	}()

	return canal_strings
}

func junta(ch3 chan string, ch1 chan string, ch2 chan string) {
	for {
		select {
		case word := <-ch1:
			ch3 <- word
		case b := <-ch2:
			ch3 <- b
		}
	}
}

func ingest(in chan string) {
	fmt.Println(<-in)
	time.Sleep(2 * time.Second)
}

func main() {
	ch1 := request_stream()
	ch2 := request_stream()
	ch3 := make(chan string)

	go ingest(ch3)
	go junta(ch3, ch1, ch2)

	ch4 := make(chan string)
	<-ch4
}
