package main

func main() {
	b := "bbb"
	println(&b, b)
	//退化赋值
	a, b := "aaa", "bbb1"
	println(&a, a, &b, b)
	//作用域
	{
		a, b := "aaa2", "bbb2"
		println(&a, a, &b, b)
	}
}
