package main

import (
	"fmt"
	"testing"
)

//定义一个map类型结构体
type TestMap map[string]string

//map结构体方法
func (t TestMap) update(k, v string) {
	t[k] = v
}

/*
	map:关联数组（哈希、字典）
*/
func Test_Map(t *testing.T) {

	//map定义
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	//map复制
	nn := n
	fmt.Println("nn: ", nn)

	//使用make创建一个map
	m := make(map[string]int)

	//设置键值对
	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 21

	//使用println打印map中所有的键值对
	fmt.Println("map", m)

	// range遍历map
	for k, v := range m {
		fmt.Println("key: ", k, "value: ", v)
	}

	//使用name[key],获取一个键的值
	v1 := m["k1"]
	fmt.Println("m['k1']=", v1)

	fmt.Println("Len: ", len(m))

	//内建的delete可以从一个map中移除键值对,删除掉map中的元素不会立马释放内存，若要释放内存只能等待指针无引用后被系统gc
	delete(m, "k2")
	fmt.Println("map: ", m)

	for k, v := range m {
		fmt.Println("key: ", k, "value: ", v)
	}

	//可选的第二返回值(如：prs，返回true/false,判断是否存在)
	//可以指出map中是否包含此键的值。避免空值0或""引起的歧义
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	//不存在，默认值为0
	fmt.Println("m['k2']: ", m["k2"])

	pr, prs := m["k1"]
	fmt.Println("pr: ", pr, "prs: ", prs)

	// 查找键值是否存在, v, ok := m["k2"] --> ok = false
	if v, ok := m["k2"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	//map 被清空,内存没有释放
	for k, _ := range n {
		delete(n, k)
	}
	fmt.Println("delete n: ", n, "nlen: ", len(n))

	//map 被清空,内存释放
	nn = nil
	fmt.Println("nnnil: ", nn, "nnnilLen: ", len(nn))

	//注意： TestMap前面没有加*，没有用指针，怎么也会改变值呢？因为map提供键值功能，用起来像指针引用的类型。类似这种功能的类型还有，数组切片，channel，interface
	var a TestMap = map[string]string{
		"a": "aa",
		"b": "bb",
	}
	fmt.Println("TestMap_a: ", a)
	a.update("a", "hello")
	a.update("b", "world")
	fmt.Println("update_TestMap_a: ", a)
}
