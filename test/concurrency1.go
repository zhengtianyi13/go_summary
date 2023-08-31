package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{} //使用waitGroup来专门控制同步
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		//wg.Add(1) //创建协程前使用，+1即计数器+1，你要是知道总共开几个协程，比如说总共开10个协程，那么直接wg.add(10)也可以，表示有10个计数器，或者说信号量
		go func(a int) {
			//wg.add(1) 常见错误，再协程里面add（1），这里要在创建协程前add，如果在协程里面add会报错
			for i := 1; i <= 100; i++ {
				fmt.Printf("%dx%d=%d\n", a, i, a*i)
			}
			wg.Done() //协程的最后需要done表示协程完成，即等于wg.add(-1)
		}(i)
	}
	wg.Wait() //wait表示等待所有协程完成,即wg中计数为0，否则阻塞

	fmt.Println("qqq")

}

//协程同步
