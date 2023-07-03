package main

import "fmt"

func main() {
	//数组
	arr := [5]int{1, 2, 3}       //指定长度
	arr1 := [...]int{1, 2, 3, 4} //不指定长度
	fmt.Println(arr, arr1)
	//切片
	slice1 := make([]int, 3)    //指定长度
	slice2 := []int{3, 1, 2}    //不指定长度
	slice3 := make([]int, 3, 5) //指定长度(len),容量(cap)
	slice4 := arr[:]            //从数组中初始化
	fmt.Println(slice1, slice2, slice3, slice4)

}
