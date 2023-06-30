package main

import (
	"errors"
	"fmt"
)

// 函数可定义多个返回值，甚至对其命名
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	} else {
		return a / b, nil
	}
}

func main() {
	//多个返回值
	a, b := 10, 2
	c, err := div(a, b)
	fmt.Println("err's value is:", err)
	fmt.Println(c, err)

	//函数是第一类型，可作为参数或返回值
	x := 100
	f := test(x)
	f()

	// 延迟调用函数
	test1(10, 1)
	allNum := addAll(1, 2, 3, 4, 5)
	println("All num is:", allNum)
}
func test(x int) func() { //返回函数类型
	return func() { //匿名函数
		println(x) //闭包
	}
}

func test1(a, b int) {
	//defer 用来释放资源、解除锁定，或执行一些清理操作,按照FILO
	// first in ，last out
	defer println("延迟调用")
	defer println("延迟调用1")
	defer println("延迟调用2")

	println(a / b)
}

/*
*
可变参函数
*/
func addAll(sum ...int) int {
	s := 0
	for e := range sum {
		s += sum[e]
	}
	return s
}
