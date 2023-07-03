package main

import "fmt"

type user2 struct {
	age  byte
	name string
}

type manager struct {
	user2 //匿名嵌入其他类型
	title string
}

func main() {
	var m manager
	m.name = "JCAK"
	m.age = 18
	m.title = "WORK"
	fmt.Println(m)
}
