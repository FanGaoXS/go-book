package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	N = 10
)

func main() {
	countCh := make(chan struct{})
	launchCh := make(chan struct{})
	exitCh := make(chan struct{})

	go func() {
		<-countCh
		log.Println("Commencing countdown")
		for i := N; i > 0; i-- {
			log.Println(i)
			select {
			case <-time.After(time.Second * 1):
				continue
			case <-exitCh:
				log.Printf("stop countdown")
				os.Exit(1)
			}
		}
		launchCh <- struct{}{}
	}()

	go func() {
		var input string
		for {
			fmt.Scan(&input)
			switch input {
			case "0":
				exitCh <- struct{}{}
			case "1":
				countCh <- struct{}{}
			}
		}
	}()

	go func() {
		<-launchCh
		log.Printf("launch")
		os.Exit(1)
	}()

	for {
	}
}
