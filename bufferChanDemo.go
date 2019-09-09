package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//有缓冲的通道和固定数目的goroutine来处理一些工作

const (
	numberGoroutines = 4  //要使用的goroutine数量
	taskLoad         = 10 //要处理的工作数量
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	//创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)
	//启动goroutine来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	//当所有工作都处理完成时关闭通道以便所有goroutine退出
	close(tasks)
	wg.Wait()
}

//worker作为goroutine启动来处理从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		//等待分配工作
		task, ok := <-tasks
		if !ok {
			//意味通道已经空了，并且已被关闭
			fmt.Printf("Worker %d : Shutting Down\n", worker)
			return
		}
		//显示我们开始工作了
		fmt.Printf("Worker %d : Started %s\n", worker, task)

		//随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		//显示我们完成工作了
		fmt.Printf("Worker %d : Completed %s\n", worker, task)
	}
}
