package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//常量不被引用也不会报错
	const x, y = 123, "yyy"
	{

	}
	//可以是编译器能计算出结果的表达式
	const (
		strSize = len("你好，go！")
		ptrSize = unsafe.Sizeof(uintptr(0))
	)
	println(ptrSize)
	println(uintptr(ptrSize))
	//在常量组中，如不指定类型和初始化值，则与上一行非空产量右值（表达式文本）相同
	const (
		a uint16 = 120
		b        //与上一行a类型、右值相同
		c = "qqq"
		d //与c类型、右值相同
	)
	fmt.Printf("%T,%v\n", b, b)
	fmt.Printf("%T,%v\n", d, d)
}
