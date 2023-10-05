package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePares(pares chan int) {
	for {
		numPar := rand.Int()
		if numPar % 2 == 0 {
			pares <- numPar
		}
	}
}

func generateImpares(impares chan int) {
	for {
		numImpar := rand.Int()
		if numImpar % 2 != 0 {
			impares <- numImpar
		}
	}
}

func consumidora(pares chan int, impares chan int) {
	for {
		select {
			case numPar, ok := <-pares:
				if ok {
					fmt.Println("par: ", numPar)
					time.Sleep(1 * time.Second)
				}
			case numImpar, ok := <-impares:
				if ok {
					fmt.Println("impar: ", numImpar)
					time.Sleep(2 & time.Second)
				}
		}
	}
}

func main() {
	pares := make(chan int)
	impares := make(chan int)
	infinito := make(chan int)

	go generatePares(pares)
	go generateImpares(impares)
	go consumidora(pares, impares)

	for {
		<-infinito
	}
}
