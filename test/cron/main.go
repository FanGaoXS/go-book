package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for _ = range ticker.C {
			fmt.Println(time.Now())
		}
	}()
	select {}
}
