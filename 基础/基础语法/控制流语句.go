package main

func main() {
	//if
	x := 100
	if x > 0 {
		println("x:", x)
	} else if x < 0 {
		println(x)
	} else {
		print(0)
	}

	//switch
	y := 100
	switch {
	case y > 0:
		println(y)
	case y < 0:
		println(y)
	default:
		println(y)
	}

	//for

	fh := ","
	for i := 0; i < 5; i++ {
		if i == 4 {
			fh = ""
		}
		print(i, fh)
	}
	println()

	//作用同while(index < 5)
	index := 0
	for index < 5 {
		print(index, ",")
		index++
	}

	println()

	//作用同while(true)
	index1 := 4
	for {
		print(index1, ",")
		index1--
		if index1 < 0 {
			break
		}
	}

	println()

	// 数组的遍历
	intArr := []int{101, 102, 103, 104, 105}
	for i, n := range intArr {
		println("the key is:", i, ",and the value is:", n)
	}
	println(intArr[1])

}
