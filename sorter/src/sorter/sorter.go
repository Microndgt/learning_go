package main

import (
	"strconv"
	"io"
	"bufio"
	"os"
	"flag"
	"fmt"
	"time"
	"algorithms/quicksort"
	"algorithms/bubblesort"
)

// 名字，默认值，使用方法
var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")


func read_values(infile string)(values []int, err error){
	file, err := os.Open(infile)
	if err != nil{
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	// 在函数return之前关闭
	defer file.Close()

	// 带缓存的IO操作
	br := bufio.NewReader(file)

	// 直接创建初始元素为0个的数组切片
	values = make([]int, 0)

	// 无限循环
	for{
		// 如果要返回的line太长，超过缓存区，那么isPrefix就是True，直到为False表示这一行已经结束
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil{
			if err1 != io.EOF{
				err = err1
			}
			// 如果有错误而且不是EOF，则会跳出循环，然后这里已经定义err，则返回的变量已经定义了
			break
		}
		if isPrefix{
			fmt.Println("A too long line, seems unexpected")
			return
		}
		fmt.Println(line)
		// 将字符数组转换为字符串，字符数组里面包含的是ASCII码
		str := string(line)
		fmt.Println(str)
		// ascii to integer
		// strconv.Itoa integer to ascii
		value, err1 := strconv.Atoi(str)
		fmt.Println(value)		
		if err1 != nil{
			err = err1
			return
		}
		
		values = append(values, value)
		fmt.Println(values)
		
	}
	return
}


func write_values(values []int, outfile string) (err error){
	file, err := os.Create(outfile)
	if err != nil{
		fmt.Println("Failed to create the output file ", outfile)
		return
	}
	defer file.Close()
	for _, value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return
}


func main(){
	flag.Parse()

	if infile != nil{
		// 输入文件不为空
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
	}

	// 读取数据
	values, err := read_values(*infile)
	if err == nil{
		fmt.Println("read values: ", values)
		t1 := time.Now()
		switch *algorithm{
		case "qsort":
			quicksort.QuickSort2(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("sorting algorithm", *algorithm, "is either unknown or unsupported")
		}
		t2 := time.Now()
		fmt.Println("sorted values:", values)
		fmt.Println("the sorting process costs", t2.Sub(t1), "to complete")
		// 写入数据
		err = write_values(values, *outfile)
		if err != nil{
			fmt.Println("write file failed ", err)
		} else {
			fmt.Println("write success")
		}
	} else {
		fmt.Println(err)
	}

	
}