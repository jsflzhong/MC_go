package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
*
关键语法:

	1.go 并发执行
	2.chan  两种操作符, chan <- 和 <- chan 分别对应往通道里放, 和从通道往外拿.

	Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

	goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

	goroutine 语法格式：
		go 函数名( 参数列表 )

	Go 允许使用 "go" 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。

	同一个程序中的所有 goroutine 共享同一个地址空间。
*/
func TestMainGoroutin1(t *testing.T) {
	// 实例化一个字符串类型的通道
	channel := make(chan string) //make ? --Refer to the statement in 'NewAndMake.go'
	// 创建producer()函数的并发goroutine
	// 并发执行一个生产者函数，两行分别创建了这个函数搭配不同参数的两个 goroutine
	go producer("cat", channel)
	go producer("dog", channel) //关于这里用的go关键字, 参考最上面注释中的说明
	// 执行消费者函数通过通道进行数据消费
	consumer(channel)
}

// 数据生产者
// 生产数据的函数，传入一个标记类型的字符串及一个只能写入的通道。
func producer(header string, channel chan<- string) {
	//无限循环, 不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		//使用 rand.Int31() 生成一个随机数，使用 fmt.Sprintf() 函数将 header 和随机数格式化为字符串。
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31()) //todo 拼接的方式?
		// 等待1秒
		//使用 time.Sleep() 函数暂停 1 秒再执行这个函数。如果在 goroutine 中执行时，暂停不会影响其他 goroutine 的执行。
		time.Sleep(time.Second)
	}
}

// 数据消费者
// 消费数据的函数，传入一个只能拿出的通道。
func consumer(channel <-chan string) { //todo <-chan --通道
	// 不停地获取数据
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Println("consumer", message)
	}
}
