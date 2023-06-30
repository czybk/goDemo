package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var a int
	// fmt.Println(a)
	// fmt.Println(&a)
	// var p *int
	// p = &a
	// *p = 111
	// fmt.Println(a)

	//字节序和指针运算

	var person *Person = new(Person)
	fmt.Println(person.age, person.name)
	fmt.Println(unsafe.Pointer(person))
	fmt.Println(uintptr(unsafe.Pointer(person)))

	i := 30
	iPtr1 := &i
	var iPtr2 *int64 = (*int64)(unsafe.Pointer(iPtr1))
	*iPtr2 = 20
	fmt.Println(i)

	//指针偏移量运算
	var person1 Person
	person1.age = 18
	person1.name = "tom"
	ptr := &person1
	fmt.Println(*ptr)
	name := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Offsetof(ptr.name)))
	*name = "汤姆"
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Offsetof(ptr.age)))
	*age = 20
	fmt.Println(*ptr)
}

type Person struct {
	age  int8
	name string
}
