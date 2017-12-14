package bubblesort

func BubbleSort(values []int){
	flag := true

	for i := 0; i < len(values) - 1; i++{
		flag = true

		for j := 0; j < len(values) - i - 1; j++{
			// 将j处的元素冒泡
			if values[j] > values[j+1]{
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		// flag为true，说明本轮没有排过序
		if flag == true{
			break
		}
	}
}


// 冒泡算法的几个要点
// 第一层循环是遍历驱动者
// 往哪边冒泡，哪边就会越来越靠近起点
// 以这种方式来确定边界和遍历起始
func BubbleSort2(values []int){
	flag := true
	for i := 0; i < len(values) - 1; i++{
		flag = true
		for j := len(values) - 1; j > i; j--{
			if values[j] < values[j-1]{
				values[j], values[j-1] = values[j-1], values[j]
				flag = false
			}
		}
		if flag{
			break
		}
	}
}


func BubbleSort3(values []int){
	flag := true
	for i := len(values) - 1; i > 0; i--{
		flag = true
		for j := 0; j < i; j++{
			if values[j] > values[j+1]{
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag{
			break
		}
	}
}