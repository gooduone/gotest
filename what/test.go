package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	signalExit := make(chan int)
	fmt.Println("begin: main")
	for i := 0; i < 2; i++ {
		go func(i int) {
			wg.Add(1)
			fmt.Println("begin: goroutine:", i)
			for {
				select {
				case <-signalExit:
					fmt.Println("end: goroutine:", i, "reason:signal")
					wg.Done()
					return
				case <-ctx.Done():
					fmt.Println("end: goroutine:", i, "reason:context", ctx.Err())
					wg.Done()
					return
				default:
					fmt.Println("work: goroutine:", i)
					time.Sleep(time.Second)
				}
			}
		}(i)

	}

	time.Sleep(time.Second * 2)
	fmt.Println("send signal: signalExit")
	//close(signalExit)
	//cancel()
	wg.Wait()
	fmt.Println("end: main")

}
