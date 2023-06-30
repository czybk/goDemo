package main

import "fmt"

// 使用指针交换
func swapByPointer(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 1, 2
	fmt.Println(&a)
	fmt.Println("before change a,b is:", a, b)

	swapByPointer(&a, &b)
	fmt.Println("after change a,b is:", a, b)

	a, b = swapByReturn(a, b)
	fmt.Println("after return a,b is:", a, b)

}

// 通过两个返回值交换
func swapByReturn(a, b int) (c, d int) {
	return b, a
}
