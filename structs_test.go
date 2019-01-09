package main

import (
	"fmt"
	"testing"
)

/*
	结构体：各个字段字段的类型的集合
*/
func Test_Structs(t *testing.T) {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	//省略的字段将被初始化为零值
	fmt.Println(person{name: "Fred"})
	//& 前缀生成一个结构体指针(?????这里为什么不直接返回地址？？？？)
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	//fmt.Println(&s)
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)
	sp.age = 51
	fmt.Println(sp.age)
}

type person struct {
	name string
	age  int
}
