package main
// 定义了包名，必须在源文件中非注释的第一行指明这个文件属于哪个包
// package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

import (
	"errors"
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

	JLoop: for j := 0; j < 5; j++{
		for i := 0; i < 10; i++{
			if i > 5{
				break JLoop
			}
		fmt.Println(i)
		}
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
	var balance_1 = []float32{1000}
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

	const MAX int = 3
	// 整形指针数组，也就是数组每一个元素的类型都是整形指针
	var ptr_array [MAX] *int
	var i int
	for i = 0; i < MAX; i++ {
		ptr_array[i] = &i
	}
	for i = 0; i < MAX; i++ {
		fmt.Println(i, *ptr_array[i])
	}

	// 指针作为函数参数
	// 允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。
	// 也就是下面的swap_ptr函数

	var Book1 Books
	Book1.author = "kevin"
	Book1.subject = "math"
	Book1.title = "Python"
	fmt.Println(Book1)

	// 结构体作为函数参数
	printBook(Book1)

	// 结构体指针
	var struct_pointer *Books
	struct_pointer = &Book1
	fmt.Println(struct_pointer.author)
	printBook_ptr(struct_pointer)

	// 切片
	// 即动态数组，切片长度不固定，可以追加元素
	// 这里赋值了一个没有设置长度的数组
	var slice_1 []int = []int{1, 2, 3}
	// 或者使用make来创建切片
	// 第二个参数是length, 第三个是capacity
	var slice_2 = make([]int, 5, 10)
	fmt.Println(slice_1)
	fmt.Println(slice_2)

	// cap()可以测量切片最长可以到达多少
	fmt.Println(cap(slice_1)) // 3
	fmt.Println(cap(slice_2)) // 10

	// nil空切片，只声明但是没有初始化前默认是nil([])，长度是0
	var slice_3 []int
	fmt.Println(slice_3)
	fmt.Println(cap(slice_3), len(slice_3)) // 0, 0

	// append()函数向切片追加新元素,变态的使用方法
	// 比起Python一点都不优雅啊
	slice_3 = append(slice_3, 5)
	slice_3 = append(slice_3, 1, 2, 3)
	fmt.Println(slice_3)
	fmt.Println(cap(slice_3), len(slice_3)) // 4, 4

	var slice_4 []int = make([]int, len(slice_3) * 2, cap(slice_3) * 2)
	copy(slice_4, slice_3)	
	fmt.Println(slice_4) // [5 1 2 3 0 0 0 0]
	fmt.Println(cap(slice_4), len(slice_4)) // 8, 8
	
	// range 范围，用于迭代数组，切片，通道和集合的元素
	// 数组和切片中返回元素的索引值，集合中返回key-value对的key值
	var sum int = 0
	for _, num := range slice_4{
		sum += num
	}
	fmt.Println(sum)
	
	for i, x := range "love"{
		fmt.Println("index:", i, "ele:", x)
	}

	// map集合
	// 第一个是key的类型，第二个是值的类型
	// 必须先声明，然后再用make初始化
	var countryCapitalMap map[string]string	
	countryCapitalMap = make(map[string]string)
	fmt.Println(countryCapitalMap)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["China"] = "Beijing"

	for country := range countryCapitalMap{
		fmt.Println("country:", country)
	}

	fmt.Println(countryCapitalMap["France"])

	// 判断key是否存在, ok是ture或者false
	capital, ok := countryCapitalMap["united states"]
	if ok {
		fmt.Println("united states is in", capital)
	}else{
		fmt.Println("united states is not in", capital)
	}

	// 另一种初始化方法
	countryCapitalMap_2 := map[string]string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
	fmt.Println(countryCapitalMap_2)

	delete(countryCapitalMap_2, "France")

	for country := range countryCapitalMap_2{
		fmt.Println(country, countryCapitalMap_2[country])
	}

	// 递归调用
	fmt.Println(fibonacci(10))

	// 接口
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// 错误处理
	// 通过实现 error 接口类型来生成错误信息
	fmt.Println(Sqrt(-1))
	fmt.Println(Sqrt(2))

	if result, errorMsg := Divide(100, 10); errorMsg == "" {
        fmt.Println("100/10 = ", result)
    }
    // 当被除数为零的时候会返回错误信息
    if _, errorMsg := Divide(100, 0); errorMsg != "" {
        fmt.Println("errorMsg is: ", errorMsg)
    }
}

func Sqrt(f float32) (float32, error){
	if f < 0{
		return 0, errors.New("math: square root of negative number")
	}
	return 1, nil
}

// 自定义一个error结构，并且实现error接口 Error()方法
type DivideError struct{
	dividee int
	divider int
}

// 实现接口Error()方法，是DivideError的
func (de DivideError) Error() string{
	strFormat := `
	Cannot proceed, the divider is zero
	dividee: %d
	divider: 0`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
        dData := DivideError{
            dividee: varDividee,
            divider: varDivider,
        }
        errorMsg = dData.Error()
        return
    } else {
        return varDividee / varDivider, ""
    }

}

// 接口
// 把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

// 这个Phone接口定义了call方法
// 因此后面的只要实现了这个方法，就是实现了这个Phone接口
type Phone interface{
	call()
}

type NokiaPhone struct{

}
func (nokiaPhone NokiaPhone) call(){
	fmt.Println("I am Nokia, I can call you")
}

type IPhone struct{

}
func (iphone IPhone) call(){
	fmt.Println("I'm IPhone, I can call you")
}


func fibonacci(n int) []int{
	var a, b int = 0, 1
	var res []int
	res = append(res, 0)
	for i := 0; i < n; i++{
		a, b = a+b, a
		fmt.Println(a)
		res = append(res, a)
	}
	return res
}

// 结构体
// 在结构体中我们可以为不同项定义不同的数据类型。
// type设定了结构体的名称，struct定义一个新的数据类型
type Books struct{
	title string
	author string
	subject string
	book_id int
}

// 传递结构体指针
func printBook_ptr(book *Books){
	fmt.Println(book.author)
}

func printBook(book Books){
	fmt.Println(book.author)
	fmt.Println(book.title)
	fmt.Println(book.subject)
	fmt.Println(book.book_id)
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