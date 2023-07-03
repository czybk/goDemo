package main

func main() {
	//引用类型特指slice、map、channel这三种预定义类型
	arr := make([]int, 0, 10) //初始化一个空切片

	//当使用append时，会检查底层数组长度是否已经不够，如果长度不够，则会新建一个数组，把原数组的数据拷贝过去，再将slice中的
	//指向数组的指针指向新的数组
	arr = append(arr, 101, 102)
	println(arr[1])

	m := make(map[string]int)
	m["key"] = 1
	println(m["key"])
	slice1 := []int{1, 2}
	println(slice1[0])
}
