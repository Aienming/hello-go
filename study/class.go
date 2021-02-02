package study

import (
	"fmt"
	"reflect"
)

/* 方法定义：
Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。
*/
// 我是这么理解的，go 中没有类的概念
// 取而代之的是，struct 结构体和 interface 接口，(不用接口也行，但是为了解耦，为了扩展？)
//				结构体里定义了“类的”属性名和类型
//				接口里定义了“类的”方法名和方法返回值得类型（或者没有返回值）
// 然后定义一个方法通过使函数名称前方 () 括号包含了接收者(结构体类型得变量)来表明这是个方法
// 语法：
// 		func (variable_name variable_data_type) function_name() [return_type]{
//    		/* 代码*/
// 		}

type action interface {
	speak() 
	Eat() string
}

type Animal struct {
	name string
	Age int
	Species string
}

func (animal Animal) speak() {
	println("my name is ", animal.name, " I'm ", animal.Age, " years old, I belong to ", animal.Species)
}

// func (animal Animal) Eat() string {
// 	var food string
// 	switch animal.Species {
// 	case "dog" :
// 		food = "shit"
// 	case "cat":
// 		food = "fish"
// 	default :
// 		food = "nothing"
// 	}
// 	println("I'm", animal.name, "I eat", food)

// 	return food
// }

// 生成狗和猫得实例
func ClassDemo() {
	var dog Animal
	var cat Animal

	dog.name = "旺财"
	dog.Age = 3
	dog.Species = "dog"

	cat.name = "mimi"
	cat.Age = 2
	cat.Species = "cat"

	dog.speak()
	// dog.Eat()	// 报错：dog.Eat undefined (type Animal has no field or method Eat)
	cat.speak()
}



// 另外一个不同的写法

type Sale interface {
	price() string
}

type Ford struct {
	brand string
}

func (f Ford) price() string {
	println("The price of Ford is $10000")
	// println(f.brand)	// 此处没有值，因为 f 只是一个 Ford 结构体的一个没有初始化的变量
	// Ford{brand: "Henry Ford"}
	// println(f.brand)
	return "$10000"
}

func (f Ford) setBrand(name string) {
	f.brand = name
}

type Chevrolet struct {
	brand string
}

func (c Chevrolet) price() string {
	println("The price of Chevrolet is $20000")
	println(c.brand)
	return "$20000"
}

func InterfaceDemo() {
	// 声明一个接口类型的变量
	var sale Sale

	sale = new(Ford)
	
	sale.price()

	// 声明一个变量的静态类型是接口，具体实现是结构体类型
	var cc Sale = Ford{"henry ford"}

	cc.price()
}

/*
	使用struct定义了数据结构，可以直接使用func方法定义数据结构中使用的方法。
	但是为了解耦，为了扩展，一般在真正设置功能性函数时，除了内置的数据类型外，都推荐使用接口的方法来传递相关方法。
*/


// 定义一个结构体（类）
type student struct {
	name string
	grade int
	score float32
}

// 通过接收者声明表示这个函数是 student 类的方法
func (s student) info() {
	println("name:", s.name)
}

// 如果要修改结构体中的属性值，就需要使用到引用传递
func (s *student) setName(name string) {
	s.name = name
}

// 实现一个study包其他文件中的接口
func (s student) learn() {
	println("student are learing")
}

func (s student) acceptInterfaceParam(i behavior) {
	i.learn()
}

// 不属于 student 结构体的方法
func setName(name string) {
	println("非任何结构体方法的函数，接收name：", name)
}

// 从这个示例可以看出 struct 几乎等于 类
// 结构体类型的方法名不能和属性重名也不能这个类型的其他方法重名
func StudentClass() {
	// 初始化这个结构体（实例化这个类？）
	student := student{"alen", 3, 99.9}

	// 获取名称
	student.info()

	// 设置名称
	student.setName("aienming")

	// 获取修改后的名字
	student.info()

	student.learn()

	// 因为 student 实现了 behavior 接口的方法，所以 student 可以作为一个接口传入参数中
	student.acceptInterfaceParam(student)

	println(reflect.TypeOf(student))
	println(fmt.Printf(`%T`, student))

	// switch student.(type) {
	// case behavior:
	// 	println("behavior")
	// case student:
	// 	println("student")
	// default:
	// 	prinln("unkonwn type")
	// }
}

