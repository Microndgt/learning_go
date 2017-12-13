package main

import (
	"fmt"
)

func my_printf(args ...interface{}){
	// 传递任意类型任意多的参数
	for index, arg := range args{
		fmt.Println("now: ", index)
		switch arg.(type){
		case int:
			fmt.Println(arg, " is an int value")
		case string:
			fmt.Println(arg, " is a string value")
		default:
			fmt.Println("can not recognize ", arg)
		}
	}
}

func myfunc(){
	i := 0
	HERE: fmt.Println(i)
	i++
	if i < 10{
		goto HERE
	}
}

type PathError struct{
	Op string
	Path string
	Err error
}

// 实现接口
func (e *PathError) Error() string{
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func main(){
	myfunc()
	recover()
	my_printf(1, 2, 3, "4")
}