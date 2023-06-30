package main

import "fmt"

func main() {
	//占位符需要用%符号表示,示例
	s := fmt.Sprintf("%s早上好", "Tom")
	fmt.Println(s)
	//普通占位符
	var example = Example{Content: "示例"}
	//%v(获取数据值)
	//%+v（获取数据值，如果是结构体，携带字段名）
	//%#v（获取数据值，如果是结构体，携带结构体名和字段名）
	fmt.Printf("%v\n", example)  //{示例}
	fmt.Printf("%+v\n", example) //{示例}
	fmt.Printf("%#v\n", example) //main.Example{Content:"示例"}

}

type Example struct {
	Content string
}
