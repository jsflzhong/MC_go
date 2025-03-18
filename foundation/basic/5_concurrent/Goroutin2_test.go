package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// 1. 使用 go 关键字启动一个新的 goroutine
func TestBasicGoroutine(t *testing.T) {
	t.Log("主线程开始")
	go func() {
		t.Log("这是一个独立的 Goroutine")
	}()
	//这行如果注释掉了,亲测则有可能下下行的"主线程结束"先执行. 就是说主线程可以先完事, 子线程可以单独运行不影响.
	time.Sleep(time.Millisecond * 10) // 确保 goroutine 有机会运行
	t.Log("主线程结束")
}

// 2. 使用 WaitGroup 确保多个 Goroutine 完成后再退出
func TestWaitGroup(t *testing.T) {
	// 创建一个 WaitGroup 变量，用于等待多个 Goroutine 结束
	var wg sync.WaitGroup
	t.Log("主线程开始")
	// 增加 WaitGroup 计数，表示有 1 个 Goroutine 需要等待
	wg.Add(1)
	go func() {
		// 当该 Goroutine 结束时，调用 Done() 让 WaitGroup 计数减 1
		defer wg.Done()
		t.Log("子 Goroutine 运行中")
	}()
	// 阻塞主 Goroutine，直到 WaitGroup 计数归零（即所有 Goroutine 执行完毕）
	wg.Wait()
	t.Log("所有 Goroutine 运行完毕")
}

// 3.1 使用带缓冲的通道同步 Goroutine
/*
	实际用途：
		主线程生产任务，子线程消费任务（典型的生产者-消费者模式）
		子线程异步处理不重要的逻辑（如日志、异步存储）
		多个 Goroutine 共享任务队列，提高并发性能
*/
func TestBufferedChannel(t *testing.T) {
	// 创建一个缓冲通道，类型为 string，容量为 1
	ch := make(chan string, 1)
	// 启动一个新的 Goroutine
	go func() {
		// 发送字符串到通道
		ch <- "Hello from Goroutine"
	}()
	// 从通道中接收(取出)数据，并赋值给 msg 变量
	msg := <-ch
	t.Log("收到消息:", msg)
}

/*
3.2
Q:那么通道可以子线程之间发送数据吗

A:
通道（channel）不仅仅是主 Goroutine 和子 Goroutine 之间的通信工具，
子 Goroutine 之间同样可以通过通道直接通信。
*/
func TestGoroutineToGoroutine(t *testing.T) {
	ch := make(chan int, 3) // 创建一个缓冲通道，最多存 3 个数

	go func() { // 生产者 Goroutine
		for i := 1; i <= 5; i++ {
			t.Logf("子线程1: 生产数据: %d", i)
			ch <- i                            // 发送数据到通道
			time.Sleep(500 * time.Millisecond) // 模拟计算时间
		}
		//注意这里,在显示的关闭通道后,下面的消费者线程就会退出.
		close(ch) // 生产结束后关闭通道
	}()

	//只要通道（channel）里有数据，Goroutine 监听通道时就会被唤醒，读取数据，
	//直到通道关闭或没有数据可读时 Goroutine 才会退出. (通道关闭就是上面的close(ch))
	go func() { // 消费者 Goroutine
		for data := range ch { // 不断从通道读取数据
			t.Logf("子线程2: 消费数据: %d", data)
		}
	}()

	time.Sleep(3 * time.Second) // 等待子线程执行完
}

/*
3.3 三个子线程,利用通道,互相合作处理数据
*/
func TestMultiGoroutineCommunication(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// 第一个子 Goroutine 发送消息
	go func() {
		ch1 <- "子 Goroutine 1 发送的消息"
	}()

	// 第二个子 Goroutine 接收 ch1 并发送到 ch2
	go func() {
		msg := <-ch1
		t.Logf("子 Goroutine 2 接收到: %s", msg)
		ch2 <- "子 Goroutine 2 处理后的消息"
	}()

	// 第三个子 Goroutine 接收 ch2 的最终数据
	go func() {
		finalMsg := <-ch2
		t.Logf("子 Goroutine 3 最终接收: %s", finalMsg)
	}()

	time.Sleep(time.Second) // 等待所有 Goroutine 完成
}

// 4. 使用无缓冲通道同步 Goroutine
/*
	缓冲通道 vs. 无缓冲通道

	所谓的无缓冲通道,就是没用第二个参数 (那参数表示最多存几个数)
	ch := make(chan string) // 无缓冲通道

	它和带缓冲通道 (make(chan string, N)) 的核心区别在于：

	无缓冲通道（阻塞模式）
		发送方 ch <- 必须等接收方 <-ch 读取后才能继续。
		用于线程间同步，确保数据立即被处理。
		作用：确保两个 Goroutine 步调一致，避免数据丢失。

	缓冲通道（异步模式）
		发送方 ch <- 可以继续执行，直到缓冲区满。
		接收方 <-ch 只有在需要时才消费数据。
		作用：适用于数据异步处理，如日志、任务队列。
*/
func TestUnbufferedChannel(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "数据到达"
	}()
	msg := <-ch
	t.Log("收到:", msg)
}

// 5. 使用 sync.Mutex 保护共享资源
// 相当于Java中的synchronized
func TestMutex(t *testing.T) {
	// 创建一个互斥锁，用于保证多个 Goroutine 访问共享资源时的互斥
	var mu sync.Mutex
	// 共享变量，下面多个 Goroutine 需要对其进行累加操作
	count := 0
	// 创建一个 WaitGroup，用于等待所有 Goroutine 执行完毕
	var wg sync.WaitGroup

	// 启动 5 个 Goroutine，每个 Goroutine 对 count 进行累加
	for i := 0; i < 5; i++ {
		// 先通知 WaitGroup 需要等待一个新的 Goroutine 结束
		wg.Add(1)
		// 注意: 每次for循环都会在这里启动一个新的 Goroutine,模仿多线程处理后面的锁和计算逻辑.
		go func() {
			// 确保 Goroutine 结束时调用 wg.Done()，减少 WaitGroup 计数
			defer wg.Done()
			// 加锁，确保只有一个 Goroutine 能修改 count
			mu.Lock()
			// 访问共享变量，进行累加
			count++
			// 解锁，允许其他 Goroutine 继续执行
			mu.Unlock()
		}()
	}
	// 阻塞主 Goroutine，等待所有 Goroutine 执行完毕
	wg.Wait()
	t.Log("最终 count:", count)
}

// 6. 使用 sync.Once 确保代码块只执行一次
/*
实际用途
	单例模式：在 Go 语言中实现单例对象初始化，防止多个 Goroutine 重复创建实例。
	懒加载：在高并发场景下，延迟初始化某个资源，确保只初始化一次，避免重复分配资源。
	配置初始化：保证某个全局配置文件或日志系统的初始化逻辑只执行一次。
	数据库连接池初始化：在并发访问数据库时，确保只创建一个连接池实例，避免多个连接池造成资源浪费。
*/
func TestSyncOnce(t *testing.T) {
	// 创建 sync.Once 变量，保证某个函数只执行一次
	var once sync.Once
	var wg sync.WaitGroup

	initFunction := func() {
		// 这个函数的逻辑只会执行一次
		// 注意, 下面开了5个线程调用这里, 但是却最终只打印了一次. 说明只执行了一次.
		t.Log("初始化只执行一次")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 调用 sync.Once 的 Do 方法，确保 initFunction 只执行一次
			once.Do(initFunction)
		}()
	}

	wg.Wait()
}

// 7. 使用 context 控制 Goroutine 生命周期
/*
实际用途
超时控制：在长时间运行的任务（如 HTTP 请求、数据库查询）中，避免任务无限等待。
协程管理：在 main 函数或某个 Handler 退出时，通知所有子 Goroutine 停止，避免资源泄露。
取消操作：在高并发场景下，可以通过 context 取消一组协程，例如多个 API 调用时，只要一个失败就取消其他请求。
*/
func TestContextCancel(t *testing.T) {
	// 创建一个可以被取消的 context
	// ####### ctx 作为子 Goroutine 监听的信号源，一旦 cancel() 被调用，所有监听 ctx.Done() 的 Goroutine 都会收到取消信号。
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			// ######### 监听 context 是否被取消
			case <-ctx.Done():
				// 一旦收到 ctx.Done() 信号，Goroutine 结束
				t.Log("Goroutine 退出")
				return
			default:
				// 如果未取消，则继续运行
				t.Log("Goroutine 运行中")
				// 模拟任务执行
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()

	// 主 Goroutine 先等待 150ms
	time.Sleep(time.Millisecond * 150)
	// ########## 然后主线程在这取消 context，通知子 Goroutine 退出.
	// 也就是触发了上面的"case <-ctx.Done():"逻辑了.
	cancel()
	// 等待子 Goroutine 退出
	wg.Wait()
	// 所有 Goroutine 结束后，主 Goroutine 退出
	t.Log("主线程结束")
}

// 8. 使用 runtime.Gosched 让出 CPU
/*
runtime.Gosched() 的作用
runtime.Gosched() 让出 CPU 资源，使其他 Goroutine 运行，但不会阻塞当前 Goroutine。
这是 Go 调度器的一部分，能让当前 Goroutine 暂停，让其他可运行的 Goroutine 先执行，之后再恢复执行。
*/
func TestGosched(t *testing.T) {
	t.Log("主线程开始")
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子 Goroutine 运行", i)
			// 让出 CPU 时间片，允许其他 Goroutine 运行
			runtime.Gosched()
		}
	}()
	time.Sleep(time.Millisecond * 50)
	t.Log("主线程结束")
}

// 9. 使用 runtime.Goexit 退出当前 Goroutine
/*
实际用途
优雅地退出 Goroutine

在某些情况下，可能希望某个 Goroutine 在特定时刻停止运行，而不影响其他 Goroutine 的执行。
比如：后台任务检测到某种终止条件后，调用 Goexit() 退出自身，而不影响主 Goroutine。
避免 Goroutine 继续执行不必要的逻辑

适用于 Goroutine 内部逻辑复杂，不想通过 return 一步步退出，而是直接终止 Goroutine。
调试和测试

在测试 Goroutine 时，可以用 Goexit() 让某个 Goroutine 强制退出，观察主 Goroutine 和其他 Goroutine 的行为。
替代方案
通常，我们不会使用 Goexit()，而是用 return 或 context.Context 进行更优雅的 Goroutine 退出管理。
例如：
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Goroutine 退出")
            return
        default:
            fmt.Println("工作中...")
            time.Sleep(time.Millisecond * 50)
        }
    }
}
这样可以让 Goroutine 基于外部信号（如超时或取消）优雅地退出，而不是直接用 Goexit() 硬性终止。
*/
func TestGoexit(t *testing.T) {
	go func() {
		// `Goexit()` 退出 Goroutine 之前，会执行这里的 `defer`
		defer fmt.Println("`Goexit()` 执行后, 会执行的反而这里的清理代码")
		fmt.Println("Goroutine 启动")
		// 立即终止当前 Goroutine，并且不会执行后续代码
		runtime.Goexit()
		fmt.Println("但是这行不会被执行")
	}()
	time.Sleep(time.Millisecond * 10)
	t.Log("主线程结束")
}

// 10. 使用 select 在多个 channel 之间选择
func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 100)
		ch1 <- "来自 ch1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 50)
		ch2 <- "来自 ch2"
	}()

	select {
	case msg := <-ch1:
		t.Log("收到:", msg)
	case msg := <-ch2:
		t.Log("收到:", msg)
	case <-time.After(time.Millisecond * 200):
		t.Log("超时")
	}
}
