package main

import (
	"fmt"
	"sync"
)

//在1和2中分别使用wg和chan的方法完成了主协程和从协程之间的同步
//现在需要完成各个协程间的同步
func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	//协程间的同步，在协程的最后需要读取管道，如果管道没有数据则阻塞
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(a int) {
			for i := a * 10; i <= a*10+10; i++ {
				fmt.Println(i)
			}
			// ch <- 1
			<-ch
			wg.Done()
		}(i)
		ch <- 1 //每个管道后需要向管道放入消息
	}

	wg.Wait() //等待子线程同步
	fmt.Println("exit")

}
