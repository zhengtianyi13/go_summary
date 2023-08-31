package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)

func t1() {

	for i := 1; i <= 10; i++ {
		ch1 <- 1
	}

}

func t2() {
	time.Sleep(100)

	for i := 1; i <= 10; i++ {
		ch2 <- 2
	}

}

func main() {
	//使用select同时监听多个管道，还可以进行超时检测
	go t1()
	go t2()
	ch := make(chan int)

	go func() {
		for {
			select {
			case one := <-ch1:
				fmt.Println(one)

			case two := <-ch2:
				fmt.Println(two)

			case <-time.After(time.Second * 3):
				fmt.Println("超时退出")
				<-ch
				return
			}
		}
	}()

	ch <- 1
	fmt.Println("退出")

}
