package study

import (
	"fmt"
	"time"
)

// 信道操作默认是阻塞的，往信道里写数据之后当前协程便阻塞，直到其他协程将数据读出。
// 一个协程被信道操作阻塞后，Go 调度器会去调用其他可用的协程，这样程序就不会一直阻塞。

/*
	关于 val, ok <- chan 的解释
	val 是接收的值，ok 标识信道是否关闭。为 true 的话，该信道还可以进行读写操作；为 false 则标识信道关闭，数据不能传输。
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func ConcurrentDemo() {
	/*
	* go 关键词开启了一个新的线程，开启的新线程就叫 goroutine
	* 意味着 go say("hello") 代码在执行时，程序不会阻塞，会立刻执行后面的 say("world")
	* 因为两个函数在两个线程里运行，所以打印的 hello 和 world 是没有固定顺序的
	*/
	go say("hello")
	say("world")
}

/*
* 通道（channel）是用来传递数据的一个数据结构。
* 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
* ch <- v    // 把 v 发送到通道 ch
* v := <-ch  // 从 ch 接收数据，并把值赋给 v
* 示例：
*                                            创建通道
*					线程 1	——————————————————————————————>
*						  /                  | |
*						/               通  | |
*		主程序 main ————                道  | |
*						\				   | |
*						 \                | |
*					线程 2  ——————————————————————————————>
*
*/

func ChannelWithBuffer() {
	// make 一个通道出来
	var ch = make(chan int, 3)

	/*
	* 向通道中添加数据
	* 没有缓冲区时，程序会阻塞等到数据被从通道中取出
	* 因为有缓冲区，且缓冲区大小为2，可以同时添加两个整型
	*/
	ch <- 1
	ch <- 2
	ch <- 3
	// 因为缓存区只能存储 3 个整型数据，当添加第 4 个时，程序会报错：fatal error: all goroutines are asleep - deadlock!
	// 线程已经被死锁了。运行时发现所有的 goroutine 都没有接收这个通道的值。
	// ch <- 4
	println("print now")

	// 通道实际是个队列，先进先出，也可以使用 range 方法配合 for 遍历通道的数据
	getter, ok := <- ch	// 非阻塞式接收，ok 为 bool 值用来判断是否接收到值
	println(getter, ok)

	// 关闭通道
	fmt.Println(ch)
	close(ch)
	fmt.Println(ch)
}

/*
* 关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入（这句话挺精辟的）
*/
func ChannelDemo() {

	// concurrentSync()			// 实现并发同步
	// notBlockGetChannelData()	// 非阻塞式获取通道数据

	// 构建一个通道 ch
	ch := make(chan int, 5)	

	// 起一个新线程运行匿名函数或调用一个函数
	go func() {
		println("并发的匿名函数")
		// 循环添加数据到通道中
		for i := 0; i < 4; i++ {
			ch <- i
			println("添加", i, "到通道中")
			time.Sleep(time.Second)	// 等上一秒，从打印结果就可以看出遍历获取数据是异步的
			println("等一秒钟")
		}
		// 数据添加完成后关闭通道，不然遍历接收时，取完数据会导致报错
		close(ch)
	}()		// 括号作用是使匿名函数被调用
	
	// 下面这行代码应该会和上面并发的匿名函数同时运行
	println("channelDemo")

	// 遍历通道中的数据，如果遍历一个未关闭的通道，数据取完后仍然在阻塞，导致 deadlock
	for data := range ch {
		println(data)
	}


	// 练习
	chP := make(chan int, 10)
	go channelPractice(chP)

	for i := 0; i < 100; i++ {
		_ = i * 100
	}

	for data := range chP {
		println(data)
	}
	
	println("主程序退出")
	
}

/*
* 在两个线程中通过 channel 事项并发同步
*/
func concurrentSync() {
	
	chSync := make(chan int)

	// 开启线程
	go func() {
		time.Sleep(time.Second * 2)
		// 向通道添加数据
		chSync <- 1

		println("添加阻塞结束")
	}()

	println("等待通道中数据")

	// 如果没有从通道里拿到数据，程序会阻塞在这里
	isok := <- chSync
	
	println("接收阻塞结束，已拿到通道中数据：", isok)
}

/*
* 非阻塞式获取通道数据
*/
func notBlockGetChannelData() {
	println("准备开始非阻塞式接收数据！！！！")

	chAsync := make(chan int)

	go func() {
		println("模拟逻辑等待5秒")
		time.Sleep(time.Second * 5)

		chAsync <- 1

	}()

	// 打印的时候会发现似乎另一个线程的匿名函数似乎并未执行
	// 因为这里使用了非阻塞方式，所以当执行 default 中的代码之后
	// 主函数执行完成了，导致匿名函数似乎没有执行
	select {
	case data := <- chAsync :
		println("接收到数据了：",data)
	default :
		println("没有接收到数据，但是我先执行到这里了")
	}

	// 这里用阻塞式等一下匿名函数的执行
	if _, ok := <- chAsync; ok {
		println("匿名函数已执行完成，退出程序")
	}

	/* 总结一下，非阻塞式的接收会导致通道中的数据永远不会被取出
	* 如果这个通道没有缓冲区的话，执行发送代码段的函数可能因为发送的数据没有被接收
	* 会导致发送数据代码后面的代码都无法得到执行
	*/
}


func channelPractice(c chan int) {
	for i := 0; i < 20; i++ {
		c <- i
	}
	close(c)
}