package main

import (
	"fmt"
)

func main() { //实现主线程等待所有子线程结束后再运行，主从之间的同步操作，但是所有子线程是并行的，其输出的数字顺序不固定

	channel := make([]chan int, 10) //新建切面，大小为10，每个元素为chan int

	for i := 1; i <= 10; i++ {

		channel[i-1] = make(chan int) //对应的切片位置创建一个管道，因为i从1开始所以i-1，对于每一个协程都创建一个通道来控制它
		go func(a int) {

			for i := 1; i <= 100; i++ {
				fmt.Printf("%dx%d=%d\n", a, i, a*i)
			}

			channel[a-1] <- 1 //错误注意这里是a不是i，这里匿名函数是一个闭包，i是一直在变的不是当前的值
		}(i)
	}

	for i, _ := range channel {
		<-channel[i]
	}

	fmt.Println("qqq")

}

//协程同步
