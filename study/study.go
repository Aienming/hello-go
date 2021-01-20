package study

import (
	"fmt"
	"encoding/json"
	"reflect"
)

func Study() {
	fmt.Println("hello Study")
}

/*********************** 数组部分 *******************************/

// 函数的参数限制了数组的长度
func withLenght(pArr [2][2]int) {
	fmt.Println(pArr[0][0])
}

// 函数的参数没有限制数组的长度
func withoutLenght(pArr [][]int) {
	fmt.Println(pArr[0][0])
}

/*
* go的数组的元素必须为同一种数据类型，声明时必须指定类型
* go的数组长度是固定的，声明时指定长度或者根据赋值时个数由go来自动推断（必须是声明时赋值才可以推断）
*/
func ArrayGrammer() {
	// 声明一个数组
	var arr [3] int
	// 声明并赋值
	var arr1 = [3]int{1,2,3}
	// 声明并赋值但不指定长度（由go来自动推断长度）
	var arr2 = [...]int{4,5,6,7}

	// 访问数组的元素，只能够通过索引来访问
	var element = arr2[0]

	fmt.Println(arr, arr1, arr2)
	fmt.Println(element)

	// 声明多维数组
	var mutiArr [3][3][5] int
	// 声明二维数组并赋值
	var doubleArr = [3][2]int{
		{1,2},
		{3,4},
		{5,6},	// 如果想要下一行单独显示 } 一维数组的最后一个元素就需要有这个逗号
	} 

	// 访问多为数组中的元素
	elementArr := doubleArr[1]
	elementArr1 := doubleArr[1][1]

	fmt.Println(mutiArr, doubleArr)
	fmt.Println(elementArr, elementArr1)

	/*
	* 调用函数是数组作为参数传递
	* 数组作为参数传递时是值传递，意思是接收数据的函数改变了数组，原数组不会改变
	* 切片作为参数传递时是指针传递，也就是地址传递，意思是，接收切片的函数改变了切片，原切片也会跟着改变
	* 切片的声明是 []int{} 中括号是不指定长度的，数组的声明是 [5]int{} 中括号是要声明数组长度的，或者使用 ... 来让go自动推断数组长度
	* 数组一定是有长度的，类似于js中的数组，这个数组包括数组中的元素、数组的长度，长度是作为数组的数据一部分存在的
	*/
	var arrDemo = [2][2]int{{1,2}, {3,4}}
	var arrDemo1 = [][]int{{5,6},{7,8,9}}	// 此种声明应该是go中的切片类型

	withLenght(arrDemo)		// 正常
	// withoutLenght(arrDemo)	// 没有限制数组长度的函数参数不接受声明了数组长度的数据

	// withLenght(arrDemo1)	// 限制了数组长度的函数参数不接受没有声明长度的数组
	withoutLenght(arrDemo1)	// 正常

}

/*********************** 数组部分结束 *******************************/


/*********************** 指针部分 **********************************/


// 变量实际上是一种占位符，用来引用计算机内存地址

/*
* 指针的定义：一个指针变量指向了一个值的内存地址。
* 在变量的类型前面使用 * 符号以声明此变量为指针变量
* 指针变量储存的是值的内存地址，在变量前加上 * 符号以访问这个指针变量的值（获取内存地址上的内容）
*/ 
func Pointer() {
	// 声明一个空指针
	var ptr *int;
	if ptr == nil { fmt.Println("ptr 是空指针") }

	// 声明一个指针变量
	// var pointer1 *int = 100	// 指针变量的值必须是内存地址
	var pointer2 *int
	var aa = 1
	// 使用 & 符号获取一个值的内存地址
	pointer2 = &aa
	fmt.Println("变量pointer2中的内存地址是：", pointer2)
	fmt.Println("变量pointer2中的内存地址存储的值是：", *pointer2)

	// 定义一个指针数组
	const M = 3
	var arr  = [3]int{1,2,3}
	var pointerArr [M]*int
	for i := 0; i < 3; i++ {
		pointerArr[i] = &arr[i]
	}
	fmt.Println(pointerArr)
	
	// for range 的用法
	for index, value := range arr {
		fmt.Println("range返回索引和对应的值：", index, value)
	}
}

/*********************** 指针部分结束 *******************************/


/*********************** 结构体部分 **********************************/
// 结构体在 go 中类似于 class 类，我理解为对象， 见 ./study/class.go
// func (s struct) functionName() {}	这个函数可以理解为 struct 的方法

// 使用 type strucr 定义结构体
type person struct {
	name string
	age int
	height int
}

// 使用结构体
func Struct() {
	andy := person{"andy", 28, 180}
	var Bob = person{}

	Bob.name = "Bob"
	Bob.age = 20
	Bob.height = 170
	fmt.Println(andy, Bob)
	fmt.Println(reflect.TypeOf(Bob))
	// 如果结构体的属性名称不是大写字母开头，那么转化json时将被忽略
	var jsonStr, err = json.Marshal(Bob)

	if err != nil {
		fmt.Println("转json出错")
		return
	}
	fmt.Printf("%T\n", jsonStr)			// 获取变量类型
	fmt.Println(reflect.TypeOf(Bob))	// 获取变量类型

	// {} 此处打印出空的对象，因为我们定义的person struct中属性不是公共的属性
	fmt.Println(string(jsonStr))

	canTransToJson()

	// 指针传值
	jack := person{name:"jack", age: 30, height: 170}
	fmt.Println("指针传值前：", jack)
	structPoiter(&jack)
	fmt.Println("指针传值后：", jack)
}
// 可以被转化为json的结构体属性名称必须大驼峰写法(首字母大写)
type canJson struct {
	Name string
	Age int
	Height int `json:"height"`	// 标记转化为json时 Height 对应的键名是 height，这种语法叫 tag
	Weight int `json:"-"`		// 标记转化为json时 Weight 被忽略
}
// 结构体转化为json
func canTransToJson() {
	var Tom canJson
	Tom.Name = "tom"
	Tom.Age = 22
	Tom.Height = 180
	Tom.Weight = 200

	fmt.Println("Tom转换前：", Tom)
	if jsonS, err := json.Marshal(Tom); err == nil {
		fmt.Println(string(jsonS))
	}
}

// 结构体指针传值
func structPoiter(person *person) {
	person.name = "allen"	// 指针传值时原结构体中的数据会被改变
}

/*********************** 结构体部分结束 *******************************/


/*********************** 切片部分 ************************************/

// 可以理解切片为动态数组，因为切片的长度是不固定的，追加元素时可使切片容量变大

func SliceDemo() {
	// 声明一个切片
	var s1 []int	// 切片不同于数组的最明显的区别就是不用声明数组长度
	var s2 = make([]int, 3, 5)	// 使用内置函数 make 声明切片，第二个参数len是切片长度，第三个参数cap是切片容量
	fmt.Println(s1, s2)

	// 切片截取
	// 格式：s[startIndex:endIndex]
	s3 := []int{1,2,3,4,5}
	get1 := s3[:]		// 获取所有
	get2 := s3[1:]		// 获取索引为 1 至最后一个元素
	get3 := s3[:2]		// 获取索引为 0 至 索引为 endIndex - 1 的元素
	get4 := s3[1:3]		// 获取索引为 1 至 索引为 endIndex - 1 的元素

	fmt.Println(get1, get2, get3, get4)

	// 内置函数 len() 可以获取切片长度，内置函数 cap() 可以获取切片容量
	s4 := make([]int, 3, 5)
	fmt.Println("切片s4的内容是：", s4)
	fmt.Println("切片s4的长度是：", len(s4))
	fmt.Println("切片s4的容量是：", cap(s4))

	// 刚声明的切片是空切片
	var s5 []int
	if s5 == nil {
		println("切片s5是空切片")	
	}

	// 内置函数 append() 可以向切片中追加元素
	s6 := make([]int, 3, 5)
	s7 := append(s6, 1)					// 函数是有返回值的
	printSlice("初始化的s6切片：", s6)
	printSlice("s7:向s6切片追加1个元素：", s7)
	s8 := append(s6, 2, 3, 4)
	printSlice("s6切片：", s6)
	// 此时的 s8 切片的cap翻倍了，因为原先的切片 s6 容量是5，append()在追加是发现容量不足会翻倍的扩充容量
	printSlice("s8:向s6切片追加3个元素：", s8)	
	s9 := append(s6, 1,2,3,4,5,6,7,8,9,10)
	// 如果一次性添加超过容量一倍的元素，返回的切片容量就不会翻倍的增加，容量会2个2个的增加
	printSlice("s9:向s6切片追加10个元素", s9)

	// 内置函数 copy() 可以复制一个切片
	// var s10 []int				// 复制不进去，因为复制到的切片没有长度和容量
	s10 := make([]int, 10, 20)		// 初始化的切片才能作为被复制到的切片，由于 s9 有13个元素，而 s10 只有10的长度，会自动舍弃十个元素之后的
	copy(s10, s9)
	printSlice("复制s9到s10", s10)
	// 正确使用 copy() 函数
	s11 := make([]int, len(s9), cap(s9))
	copy(s11, s9)
	printSlice("正确的复制s9切片：", s11)


	// 基于数组或切片生成新的切片
	var s12 = [8]int{1,2,3,4,5,6,7,8}
	s13 := s12[1:5]
	fmt.Println("s12数组：" ,"len=",len(s12), "cap=", cap(s12), "slice=", s12)
	// 新生成的切片容量只有7，原数组容量为8，是因为新切片是基于原数组索引 1 的位置截取的，容量会从索引 1 开始直到继承数组末尾
	printSlice("基于s12数组生成的切片s13：", s13)	// len=4, cap=7, slice=[2 3 4 5]
	// 如果此使改变 s13 切片中的某个值，原数组 s12 也会发生相应改变，说明新的切片基于原数组的内存地址生成，属于引用赋值
	s13[1] = 999
	fmt.Println(s12, s13)
	
	/*
	* 引用赋值是指在一个新的内存地址保存了原变量的内存地址。意味着引用赋值生成的变量一旦发生改变，原变量也会相应改变
	* 值传递是指将原变量的值复制一遍存在新的内存地址中
	* 指针传递指的是传递直接是原变量的内存地址，相比引用传递，引用传递多了个中间内存地址
	*/


	// 利用 append() 合并多个一维数组
	s14, s15, s16 := []int{1,2,3}, []int{1,2,3}, []int{1,2,3}
	s17 := append(append(s14, s15...), s16...)	// 这个 ... 是什么意思我还不知道
	printSlice("s14:", s14)
	printSlice("s15:", s15)
	printSlice("s16:", s16)
	printSlice("s14+s15+s16合并后s17:", s17)

	/*
	* 切片在函数传参的时候是引用传递
	* 也就是说，如果切片在调用的函数中被修改过，那么原切片也会发生相应的变化
	* 类似传递数组时使用指针传递（取数组的内存地址传递）
	*/

}

func printSlice(info string, x []int){
	fmt.Println(info,"len=",len(x), "cap=", cap(x), "slice=", x)
 }

/*********************** 切片部分结束 **********************************/


/*********************** 范围 *****************************************/

// 解释：Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
// 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。 

func RangeDemo() {
	// range 肯定不能单独存在，必须要搭配 for 循环使用。可以理解为 for + range = foreach
	strs := []string{"alen", "bob", "charlie", "eliana"}

	for val, key := range(strs) {
		println("key:", key, "val:", val)
	}
}

/*********************** 范围结束 **************************************/


/*********************** 集合 *****************************************/

// Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
// Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。 

func MapDemo() {
	// 声明一个集合
	var map1 map[int]int				// 声明但并未初始化
	map2 := make(map[string]string)		// 声明并初始化集合，此使集合默认为 nil

	fmt.Println(map1, map2)

	// 声明并初始化一个集合
	countryMap := map[string]string{"asia":"china", "america":"america", "europe":"russia"}
	// 访问集合
	asiaCountry := countryMap["asia"]
	fmt.Println(asiaCountry)	// "china"
	// 添加元素
	countryMap["africa"] = "rwanda"
	fmt.Println(countryMap)		// map[africa:rwanda america:america asia:china europe:russia]
	// 遍历集合
	for index, value := range(countryMap) {
		fmt.Println(value, "是", index, "的国家")
	}
	/*
	* 因为是无序集合，因此上面的遍历每次结果输出的顺序都是不一样的
	*/

	// 判断某个 key-value 对是否存在
	// continet := "south america"
	continet := "asia"
	value, res := countryMap[continet]
	// res 会返回 bool 值，true || false
	if res {
		println(continet, "存在, 国家有：", value)
	} else {
		println(continet, "不存在")
	}

	// 删除集合中的元素
	delete(countryMap, "africa")
	fmt.Println("删除非洲键值对后：", countryMap)

}

/*********************** 集合结束 **************************************/


/*********************** 接口 ****************************************/

// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。




/*********************** 接口结束 ************************************/

// 见 ./study/class.go

/*********************** 其他 ****************************************/

func TypeTranslate() {
	var a int32 = 10
	var b int64
	b = int64(a)	// 显式类型转换，因为变量 b 已被声明为 int64 类型，此使 b = a 会报错，因为 a 是 int32 类型
	println(a, b)
}
