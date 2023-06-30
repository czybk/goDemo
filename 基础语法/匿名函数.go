package main

import "fmt"

type user1 struct {
	name string
	age  byte
}

func (u user1) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager1 struct {
	user1
	title string
}

func main() {
	var m manager1
	m.name = "Tom"
	m.age = 40
	println(m.ToString())
}
