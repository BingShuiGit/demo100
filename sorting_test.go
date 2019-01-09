package main

import (
	"fmt"
	"sort"
	"testing"
)

//创建一个为内置 []string 类型的别名的ByLength 类型
type ByLength []string

//实现了 sort.Interface 的 Len，Less和 Swap 方法，这样我们就可以使用 sort 包的通用Sort 方法了
//Len 和 Swap 通常在各个类型中都差不多，Less 将控制实际的自定义排序逻辑
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func Test_Sorting(t *testing.T) {
	strs_Sorting := []string{"c", "a", "b"}
	sort.Strings(strs_Sorting)
	fmt.Println("strs_Sorting: ", strs_Sorting)

	ints_Sorting := []int{7, 2, 4}
	sort.Ints(ints_Sorting)
	fmt.Println("ints_Sorting: ", ints_Sorting)

	//可以使用 sort 来检查一个序列是不是已经是排好序的
	s_Sorting := sort.IntsAreSorted(ints_Sorting)
	fmt.Println(s_Sorting)

	//自定义排序
	fruits := []string{"peach", "banana", "kiwi"}
	//将原始的 fruits 切片转型成 ByLength 来实现我们的自定义排序
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
