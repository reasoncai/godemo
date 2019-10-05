package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}

		}
	}()

	var sum int
	for e := range intChan {
		fmt.Printf("received:%v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("got:%v\n", sum)
			break
		}
	}
	fmt.Printf("end\n")
}
