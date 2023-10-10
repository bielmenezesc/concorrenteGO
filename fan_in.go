package main

import (
	"fmt"
)

func request_stream() chan string {
	canal_strings := make(chan string)

	for i := 0; i < 100; i++ {
		canal_strings <- "abacate"
	}

	return canal_strings
}

func junta_canais(ch1 chan string, ch2 chan string, ch3 chan string) chan string {
	for i := 0; i< 200; i++ {
		select {
			case word, ok := <-ch1:
				if ok {
					ch3 <- word
				}
			case word2, ok := <-ch2:
				if ok {
					ch3 <- word2
				}
		}
	}

	return ch3
}

func ingest(in chan string) {
	
}

func main() {
	ch1 := request_stream()
	ch2 := request_stream()
	ch3 := make(chan string)
	

	go ingest()
}
