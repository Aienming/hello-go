package study

import (
	"fmt"
)

// 此小样照抄自： https://zhuanlan.zhihu.com/p/63354092
func DeferDemo() {
	fmt.Println(d1(1))	// 1
	fmt.Println(d2(1))	// 2
	fmt.Println(d3(1))	// 1
	fmt.Println(d4(1))	// 1

	/*
	defer 关键词多用于关闭文件句柄、释放资源时用，采用后进先执行的规则
	例：在 f := os.open() 打开文件后，确认文件读取成功后可以直接 deferf.close() 关闭文件资源，
		防止文件使用完后忘记关闭文件浪费系统资源啊
		但是 defer 开销比在结束时调用函数关闭资源 大
	还有一种在不知道何时程序会报错时提前声明 defer 语句，并配合 recover() 拦截报错信息时使用（我觉得类似于其他语言的try{}catch(){}语法
	例：func b() {
		defet func() {
			if i := recover(); i != nil {
				println("报错信息是：", i)
			}
			// do something
			panic("假装报错")
			// do otherthing
		}()
	}
	*/
}

// p return后，defer 后的匿名函数执行 p++ ，所以打印的是传进来的值
func d1(p int) int {
	defer func() {
		p++
	}()
	
	return p
}

// p return后，defer 后的匿名函数执行 r++ , 
// 因为 r 是返回值，匿名函数使用函数外的值是 引用传递，
// 所以 p 虽然以原值返回了，但由于引用的特性，匿名函数改变了 r 值，返回的值也就被篡改了
// 所以打印的是改变后的值
func d2(p int) (r int) {
	defer func() {
		r++
	}()

	return p
}

// p return后，defer 后的匿名函数执行的 p++
// 因为匿名函数和的 p++ 是通过匿名函数的参数传递的，go 中函数参数的传递都是值传递
// 所以匿名函数中的 p 值是值传递，相当于是个新的变量，因此不会改变匿名函数外的变量值
// 所以打印的是传进来的原值
func d3(p int) (r int) {
	defer func(p int) {
		p++
	}(p)

	return p
}

// p return后，defer 后的匿名函数执行的 r++
// 原因同 d3() ，匿名函数的 r 属于值传递，改变时不影响返回值 r 的值
func d4(p int) (r int) {
	defer func(r int) {
		r++
	}(r)

	return p
}
