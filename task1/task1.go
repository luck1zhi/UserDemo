package main

import "fmt"

func main() {

	slice := getSlice()
	if slice == nil{
		return
	}
	fmt.Println("单次出现的值：", getSingle(slice))
}

func getSlice() []int{

	fmt.Println("数组长度：")
	var length int = 0
	fmt.Scanln(&length)
	if length%2 == 0{
		fmt.Println("设置有误")
		return nil
	}
	slice := make([]int,length)
	fmt.Println("输入数组值：")
	for i:=0;i<length;i++{
		fmt.Scan(&slice[i])
	}
	return slice
}

func getSingle(slice []int) int {
	//结果默认为最后一个
	result := slice[len(slice) - 1]
	for i:=0;i < len(slice) - 1;i=i+2{
		if slice[i]==slice[i+1] {
			continue
		}else{
			//只可能出现在左边
			result = slice[i]
			break
		}
	}
	return result
}

/**
输出样例：

	数组长度：
	1
	输入数组值：
	1
	单次出现的值： 1
---------------------
	数组长度：
	5
	输入数组值：
	-1 -1 2 3 3
	单次出现的值： 2
---------------------
	数组长度：
	6
	设置有误



 */
