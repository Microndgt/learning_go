package quicksort


func quick_sort(values []int, ele int)(sorted_values []int){
	if len(values) > 0{
		left_values := make([]int, 0)
		right_values := make([]int, 0)
		for _, value := range values{
			if value < ele{
				left_values = append(left_values, value)
			} else {
				right_values = append(right_values, value)
			}
		}
		sorted_left_values := make([]int, 0)
		if len(left_values) > 0{
			sorted_left_values = quick_sort(left_values[1:], left_values[0])
		}

		sorted_right_values := make([]int, 0)
		if len(right_values) > 0{
			sorted_right_values = quick_sort(right_values[1:], right_values[0])
		}
		
		sorted_values = append(sorted_values, sorted_left_values...)
		sorted_values = append(sorted_values, ele)
		sorted_values = append(sorted_values, sorted_right_values...)
	} else {
		sorted_values = append(values, ele)
	}
	return sorted_values
}

func QuickSort(values []int)(sorted_values []int){
	ele := values[0]
	return quick_sort(values[1:], ele)
}

func quick_sort2(values []int, left, right int){
	temp := values[left]
	p := left
	i, j := left, right
	// 原地将整个元素切分成切分元素两边有序的
	for i <= j{
		// 一直迭代到有元素小于 切分元素
		for j >= p && values[j] >= temp{
			j--
		}
		if j >= p{
			values[p] = values[j]
			p = j
		}
		
		// 查找下一个元素是不是大于切分元素
		if values[i] <= temp && i <= p{
			i++
		}

		if i <= p{
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p - left > 1{
		quick_sort2(values, left, p-1)
	}
	if right - p > 1{
		quick_sort2(values, p+1, right)
	}
}

func QuickSort2(values []int){
	quick_sort2(values, 0, len(values)-1)
}