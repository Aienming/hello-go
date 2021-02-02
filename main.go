package main

import "fmt"
import "unsafe"
import "./study"

type MY int		// 声明MY类型是int类型，意味着MY和int有着同样的意思了


var g = "全局变量g"	// 全局变量必须用关键词声明

/*
* 关于变量作用域，函数体和循环体中的变量为局部变量
* 如果局部变量和全局变量名字相同，会优先使用局部bianliang
* 在for循环条件中声明的控制变量只会在循环体中生效，如果这个控制变量和局部变量相同，会优先使用这个控制变量
* go中不存在php中那种for/foreach的控制变量/循环变量溢出的问题，控制变量和循环变量值在循环体内能被获取
*/
func varibleScope() {
	// 局部变量
	var g = "局部变量g";

	fmt.Println("局部变量g取代了同名的全局变量g:", g)

	for g := 300; g < 302; g++ {
		fmt.Println("循环体中的控制变量g取代了局部变量g：", g)
	}

	fmt.Println("控制变量g不会溢出循环体，所以循环体外g是局部变量：", g)
	
}

// 大括号不可位于一行的开头
func circleDemo() {
	fmt.Println("go的循环关键词只有for一个:")
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("continue和break可以用label跳转到指定的循环体：")
	out: 
		for i := 1; i <= 3; i++ {
			fmt.Println(i)
			for j := 10; j <= 13; j++ {
				fmt.Println(j)
				break out;	// 可以理解为，将break放到了out这个labei代表的循环体里去执行去了
			}
		}

}


func main() {

	println("HELLO")
	fmt.Println("HELLO")

	// circleDemo()
	// varibleScope()

	// study.ArrayGrammer()			// 数组语法
	// study.Pointer()				// 指针
	// study.Struct()				// 结构体
	// study.SliceDemo()			// 切片
	// study.RangeDemo()			// 范围
	// study.MapDemo()				// 集合
	// study.ConcurrentDemo()		// 并发
	// study.ChannelWithBuffer()	// 有缓冲区的通道
	// study.ChannelDemo()			// 通道

	// study.DeferDemo()			// defer例子（defer后的语句在函数结束或发生 panic 后执行，使用后进先出的方式执行）
	// study.ClassDemo()			// 模拟一个类
	study.InterfaceDemo()		// interface 的用法

	// study.FakeMain()
	// study.StudentClass()
	


	return 
	/* := 只有在函数体中才可以使用
	* 此为初始声明符号
	* 此符号左侧不得出现已声明的变量
	* 此符号可以简写声明语句
	*/
	b := 2	// 编译器会根据值推断变量类型

	var c int = 3	// 手动声明变量类型

	// & 可以获得变量的内存地址
	e := &c

	b, c = c, b	// 快速交换两个变量的值，交换的两个变量的类型必须是相同的
	
	var _, f = 10, 11	// _ 用于抛弃值，用在函数返回多个数据时，有不需要的数据(go中声明的变量必须被使用到)

	g, h, i := 7, 8, 9	// 多个变量声明，会进行类型推断

	var a = 1
	a = a * 10
	println(a, b, c, e, f, g, h, i)

	// 常量只可以是布尔型、整型（浮点）、字符串类型
	const HEIGHT = 180

	var str = "popular"
	println(len(str), unsafe.Sizeof(str))

	// 二进制
	var bin uint = 60
	var bin2 uint = 13

	var rr uint = bin & bin2
	fmt.Printf("&结果： %d", rr)
	
	// 下面是一个判断偶数的实例
	// var s int ;    // 声明变量 s 是需要判断的数
    // fmt.Println("输入一个数字：")
    // fmt.Scan(&s)

    // if s%2 == 0  { //     取 s 处以 2 的余数是否等于 0
    //     fmt.Print(s, "是偶数\n") ;//如果成立
    // }else {
    //     fmt.Print("s 不是偶数\n") ;//否则
    // }
	// fmt.Print("s 的值是：",s) ;
	

	if only := 1; only < 10 {
		// 条件表达式中声明的变量只有这个条件语句代码块能够使用
	}


	/* 
	* switch中需要判断的值可以在switch关键词后面，也可以在每个case关键词携程表达式
	* go中switch的case语句默认结尾加有break，所以不需要手动写出来，
	* 当case后面不是表达式时可以匹配多个值，值之间用逗号隔开
	* 被比较的值和case中的值应当是同一类型
	*/
	var grade string = "B"
	var marks int = 90
 
	switch marks {
	   case 90: grade = "A"
	   case 80: grade = "B"
	   case 50,60,70 : grade = "C"
	   default: grade = "D"  
	}
 
	switch {
	   case grade == "A" :
		  fmt.Printf("优秀!\n" )    
	   case grade == "B", grade == "C" :
		  fmt.Printf("良好\n" )      
	   case grade == "D" :
		  fmt.Printf("及格\n" )      
	   case grade == "F":
		  fmt.Printf("不及格\n" )
	   default:
		  fmt.Printf("差\n" );
	}
	fmt.Printf("你的等级是 %s\n", grade );  


	/* 添加fallthroug关键词可以忽略case后的break，执行下一个case
	* 如果想要所有的case都执行，那么所有的case后都要加上fallthrough关键词，不然为遇到隐藏的break关键词，导致退出了switch语句
	* 也可以手动添加break关键词主动终止case中代码的执行
	*/
	switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case true:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
    }
}
