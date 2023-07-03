package main

func main() {
	//相同类型的多个变量,类型定义在变量后
	var (
		x, y int
		a, s = 100, "abc"
	)
	x = 100
	y = 50
	println(x)
	println(y)
	println(a)
	println(s)
}
