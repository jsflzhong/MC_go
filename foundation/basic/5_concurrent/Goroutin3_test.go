package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
*
关于select的用法

select 只处理:通道通信。
监听多个通道，只执行其中一个可用的 case（如果多个 case 可用，随机选择一个）。
如果所有通道都没数据，会阻塞，除非有 default。
*/

// case1: 监听多个通道
func TestSelect1(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 模拟异步任务
	go func() {
		time.Sleep(time.Millisecond * 100)
		ch1 <- "任务 1 完成"
	}()
	go func() {
		time.Sleep(time.Millisecond * 50)
		ch2 <- "任务 2 完成"
	}()

	// `select` 监听多个通道
	select {
	case msg1 := <-ch1:
		fmt.Println("收到:", msg1)
	case msg2 := <-ch2:
		fmt.Println("收到:", msg2)
	default:
		fmt.Println("没有数据可用，执行默认操作")
	}
}

// case2: 设置超时
func TestSelect2(t *testing.T) {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("收到:", msg)
	//如果 ch 在 100ms 内没有数据，程序不会一直阻塞，而是执行 time.After 这个 case：
	case <-time.After(time.Millisecond * 100): // 超时
		fmt.Println("超时了")
	}
}

// case3: 退出 Goroutine
func TestSelect3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done(): // 监听 `ctx` 取消信号
				fmt.Println("Goroutine 退出")
				return
			default:
				fmt.Println("Goroutine 运行中")
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()

	time.Sleep(time.Millisecond * 150)
	cancel() // 触发 `ctx.Done()`
	time.Sleep(time.Millisecond * 50)
}
