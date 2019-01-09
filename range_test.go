package main

import (
	"fmt"
	libf "libfunc"
	"testing"
)

/*
*range遍历：range可以迭代各种各样的数据结构
 */
func Test_Range(t *testing.T) {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", sum)

	for i, num := range nums {
		if num == 3 {
			prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", i)
		}
	}

	//遍历修改map的值
	m := map[string]int{"k1": 1, "k2": 2, "k3": 3}
	for kk, vv := range m {
		m[kk] = vv + 4
	}
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", m)

	//遍历打印map键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//返回索引和Unicode编码
	for i, c := range "go" {
		prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", i, c)
	}

	//for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，就会导致错误
	sli := []int{0, 1, 2, 3}
	fmt.Printf("sliaddr: %p\n", sli)
	myMap := make(map[int]*int)

	for index, value := range sli {
		//遍历过程中创建了一个新的地址存放value（每个元素的副本，遍历结束时，value的地址存放数据3，而遍历过程中，value的地址一值没有改变，所以myMap中存放的地址都是value的地址，值为3）
		fmt.Printf("valueaddr: %v value: %v\n", &value, value)
		myMap[index] = &value
		//		fmt.Printf("myMap: %v\n", myMap)
	}
	fmt.Println("=====new map1=====")
	prtMap(myMap)

	//修正错误
	for index, value := range sli {
		//每次遍历都会新建一个地址来存放value
		num := value
		//将新地址放入map中
		fmt.Printf("numaddr: %v\n", &num)
		myMap[index] = &num
	}
	fmt.Println("=====new map2=====")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
