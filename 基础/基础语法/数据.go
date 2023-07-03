package main

import "fmt"

func main() {
	//slice切片
	x := make([]int, 0, 5) //创建容量为5的切片
	for i := 0; i < 8; i++ {
		x = append(x, i) //追加数据。当超出容量限制时，自动分配更大的存储空间
	}
	fmt.Println(x)

	//字典类型map对象
	m := make(map[string]int)
	m["a"] = 1
	x1, res := m["b"]
	x2, res1 := m["a"]
	if res1 {
		fmt.Println("ok, x1 has value,and It's value is:", x2)
	} else {
		fmt.Println("oh,my god,the value is null")
	}

	// x2 = m["c"]
	fmt.Println(x1, res)
	fmt.Println(x2, res1)

	x3, res2 := m["a"]
	fmt.Println(x2, res1)
	if !res2 {
		fmt.Println("不能取到字典对象key为a的值")
	}
	m["d"] = 2
	x3, res2 = m["d"]
	fmt.Println(x3, res2)

}
