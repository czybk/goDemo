package main

import "fmt"

func main() {
	//切片
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
	//字典
	m := make(map[string]int)
	m["a"] = 1
	x1, ok1 := m["b"]
	fmt.Println(x1, ok1)
	delete(m, "a")

	var manager manager
	manager.age = 40
	manager.name = "Tom"
	manager.title = "COO"
	fmt.Println(manager)
}

/*
*
结构体
*/
type user struct {
	name string
	age  byte
}

/**
*匿名嵌入其他类型
 */
type manager struct {
	user
	title string
}
