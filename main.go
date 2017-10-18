package main
// 定义了包名，必须在源文件中非注释的第一行指明这个文件属于哪个包
// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

import (
	"fmt"
	"math"
)

// func main() 是程序开始执行的函数
// main 函数是每一个可执行程序所必须包含的
// 一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
func main(){
	// 如果标识符以大写字母开头，那么这种形式的标识符代表的对象就可以被外部包的代码使用，小写字母开头，则对外部包是不可见的
	fmt.Println("hello world")

	// 变量声明
	// go的字符串是由单个字节连接起来的
	// 字节使用utf-8编码标识Unicode文本
	// 类型在变量后面声明
	// 字符串类型在go里是个结构，包含指向底层数组的指针和长度，每个部分都是8个字节
	var name string
	name = "kevin"
	fmt.Println(name)

	var male bool = true;
	fmt.Println(male)

	// 字符串的默认值是空
	var default_ string
	fmt.Println(default_)

	// 自动判定变量类型，也不能声明以前声明过的
	var string_ = "love"
	fmt.Println(string_)

	// 省略var,但是:=左侧的变量不应该是声明过的
	// 只能在函数体中出现
	test_ := "test_1"
	fmt.Println(test_)

	// 声明多个变量
	var vname1, vname2, vname3 string
	var vname4, vname5, vname6 = test_, default_, male
	fmt.Println(vname1, vname2, vname3, vname4, vname5, vname6)

	// 全局变量
	var (
		vname7 string
		vname8 string
	)
	fmt.Println(vname7, vname8)

	// 看起来go的数据类型和Python基本一致，值类型和引用类型
	// := 使用这个只能在函数体内声明新变量，初始化声明
	// 声明只有一次，后面对这个值的类型将不能被变化
	// 局部变量声明就必须被使用
	var a = 30
	fmt.Println(a)

	// 全局变量可以允许声明但不使用
	// 可以和Python一样这样使用
	var b = 40
	a, b = b, a

	// 常量
	const f string = "abc"
	const c = "def"
	const d, e = f, c
	fmt.Println(f, c, d, e)
	// 常量用作枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
	fmt.Println(a, b, c)
	
	// & 返回变量存储地址
	// * 指针变量
	var ptr *int
	var ptr_int int = 3
	ptr = &ptr_int
	fmt.Println("ptr_int's value is", ptr_int)
	fmt.Println("ptr's value is", *ptr)
	fmt.Println("ptr is", ptr)

	// 条件语句
	if ptr_int < 5 {
		fmt.Println("ptr_int is less than 5")
	}else{
		fmt.Println("ptr_int is more than 5")
	}

	// switch语句
	switch ptr_int {
		// case后面的变量必须是相同类型
		// 也可以这样 case val1, val2, val3
	case 1:
		fmt.Println("it's 1")
	case 2:
		fmt.Println("it's 2")
	case 3, 4:
		fmt.Println("3, 4")
	default:
		fmt.Println("it's", ptr_int)
	}

	// 用于type-switch来判断interface变量中实际存储的变量类型
	// interface是一组method的组合，我们通过interface来定义对象的一组行为。
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("x type is", i)
	case int:
		fmt.Println("x type is int")
	case float64:
		fmt.Printf("x 是 float64 型")           
	 case func(int) float64:
		fmt.Printf("x 是 func(int) 型")      
	 case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )  
	 default:
		fmt.Printf("未知型") 
	}

	// 循环
	var loop_a, loop_b int = 15, 25
	numbers := [6]int{1, 2, 3, 5}
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	for i, x := range numbers {
		fmt.Println("index", i, "number", x)
	}
	for loop_a < loop_b {
		loop_a++
		fmt.Println(loop_a)
	}

	// goto语句
	// goto语句需要对要跳转的语句进行标记
	var goto_a = 10
	LOOP: for goto_a < 20{
		if goto_a == 15 {
			goto_a++
			// 相当于跳过15的打印
			goto LOOP
		}
		fmt.Println("goto_a:",goto_a)
		goto_a++
	}

	// 函数
	// 并不能在函数内部定义函数
	var res int
	var func_2 int = 5
	res = max('b', func_2)
	fmt.Println(res)

	func_3, func_4 := swap("a", "b")
	fmt.Println(func_3, func_4)

	// 传递引用
	// 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	ptr_1, ptr_2 := 100, 200
	swap_ptr(&ptr_1, &ptr_2)
	fmt.Println(ptr_1, ptr_2)

	// 函数是一级对象，Python中的，Go中函数也可以作为值使用
	func_5 := math.Sqrt
	fmt.Println(func_5(64.0))

	// 闭包，可以直接使用函数内部的变量，不必声明
	// 和Python中用法是一致的
	nextNumber := getSequence()
	fmt.Println(nextNumber()) // 1
	fmt.Println(nextNumber()) // 2
	fmt.Println(nextNumber()) // 3

	// 方法
	// 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
	// 就是如果给定了类型，就是这个类型的对象的方法

	var c1 Circle
	c1.radius = 10
	fmt.Println("radius:", c1.radius)
	fmt.Println("area", c1.getArea())

	// 数组
	var balance [10] float32
	balance[0] = 1
	// 数组初始化,未初始化的元素就是0
	var balance_0 = [5]float32{1000}
	fmt.Println(balance_0)
	// 可以根据元素个数设置数组大小,这个数组大小为1
	var balance_1 = [...]float32{1000}
	fmt.Println(len(balance_1))

	// 多维数组
	var balance_5 = [3][4] int {
		{0, 1, 2, 3},
		{1, 2, 3, 4},
	}
	fmt.Println(balance_5)

	// 指针
	// int类型的指针
	var point_1 *int
	var point_2 int = 10
	point_1 = &point_2
	fmt.Println("point_2:", point_2)
	fmt.Println("point_1:", *point_1)
	*point_1 = 30
	fmt.Println("point_2:", point_2)
	fmt.Println("point_1:", *point_1)	
	point_2 = 50
	fmt.Println("point_2:", point_2)
	fmt.Println("point_1:", *point_1)

	// 空指针
	// 值为nil，比较的话 ptr != nil
	var ptr_nil *int
	fmt.Println(ptr_nil)
	
}

// 定义结构体类型
type Circle struct{
	radius float32
}

// 该方法属于Circle结构体类型对象中的方法
func (c Circle) getArea() float32{
	// radius即为Circle类型对象的属性
	return 3.14 * c.radius * c.radius
}

func getSequence() func() int {
	i := 0
	return func() int{
		i += 1
		return i
	}
}

// 值传递
func swap(x string, y string) (string, string){
	return y, x
}

// 引用传递
func swap_ptr(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func max(num1 int, num2 int) int {
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
	var result int
	if (num1 > num2) {
		result = num1
	}else{
		result = num2
	}
	return result
}