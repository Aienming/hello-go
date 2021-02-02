package study

import (
	// "fmt"
	"time"
)

// class.go 有用
type behavior interface {
	learn()
}


func thread1(ch chan int) {
	for i := 1; i <= 10; i++ {
		if(i == 5) {
			println("通道堵塞中")
		}
		ch <- i
	}
	println("通道数据插入完毕，关闭通道")
	close(ch)
}

func thread2() {
	println("线程二运行中..")
	time.Sleep(time.Second * 1)
	println("线程二运行中")
	time.Sleep(time.Second * 1)
	println("线程二运行中")
	time.Sleep(time.Second * 1)
	println("线程二运行中")
	time.Sleep(time.Second * 1)
	println("线程二运行结束")
}

func thread3(ch chan int) {
	// println("获取通道种的数据")

	
	// println("线程三完成")
}



// 开启三个线程
func FakeMain() {
	ch := make(chan int, 5)
	println("开始")
	println("开启线程一")
	go thread1(ch)
	println("开启线程二")
	go thread2()
	// println("开启线程三")
	// go thread3(ch)

	println("主线程休息五秒")
	time.Sleep(time.Second * 5)

	for {
		var val, ok = <- ch
		if(!ok) {
			println("通道被关闭")
			break;
		}
		println("取出数据：", val)
	}

	println("主程序完成")

}
