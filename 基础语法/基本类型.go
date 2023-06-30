package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//基本类型
	a, b, c := 100, 0144, 0x64
	fmt.Println(a, b, c)
	//%b二进制、%#o八进制（输出带0开头的八进制 ）、%#x 输出带 0x 开头的十六进制
	fmt.Printf("0b%b,%#o,%#x\n", a, a, a)
	fmt.Println(math.MinInt8, math.MaxInt8)
	//base为输入的数据类型
	d, _ := strconv.ParseInt("1000110", 2, 32)
	println(d)
	ccc := "222"
	println(ccc)

}
