package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("now:%v.\n", time.Now())
	//这里会阻塞3秒
	expirationTime := <-timer.C
	fmt.Printf("expiration time:%v.\n", expirationTime)
	fmt.Printf("Stop timer:%v.\n", timer.Stop())
}
