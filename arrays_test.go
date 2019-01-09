package main

import (
	"fmt"
	libf "libfunc"
	"testing"
)

/*
	数组: 数组需要指定大小，不指定也会根据初始化的自动推算出大小，不可改变
*/
func Test_Arrays(t *testing.T) {
	var a [5]int
	//默认int型数组初始值是0，string型为空，“ ”
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("Set:", a)
	fmt.Println("Get:", a[4])
	fmt.Println("Len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "dcl:", b)

	//以下两行证明数组是值传递
	change_Array(b)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "b: ", b)

	s := []int{1, 2, 3, 4}
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "s: ", s)
	//以下两行证明切片是引用传递
	change_Slice(s)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "s: ", s)

	c := [...]int{1: 1, 0: 2}
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "c: ", c)

	var aa [2]string
	fmt.Println("aa: ", aa)

	var ab [2]int = [2]int{}
	fmt.Println("ab: ", ab)

	var bc = [2]int{1, 2}
	fmt.Println("bc: ", bc)

	cc := [3]int{}
	fmt.Println("cc: ", cc)

	bb := [3]int{1, 2}
	fmt.Println("bb: ", bb)

	e := [5]int{4: 1}
	fmt.Println("e: ", e)

	//不指定也会根据初始化的自动推算出大小
	gg := [...]int{}
	fmt.Println("gg: ", gg)
	fmt.Println(typeof(gg))

	g := [...]int{4: 1}
	fmt.Println("g: ", g)

	f := [...]int{4, 1, 3}
	fmt.Println("f: ", f)

	h := [...]int{0: 4, 1: 1, 2: 3}
	fmt.Println("h: ", h)

	hh := [...]int{0: 4, 1, 2}
	fmt.Println("hh: ", hh)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Printf("2D: %v\n", twoD)

	//数组指针和指针数组：注意*与谁结合，如p *[5]int，*与数组结合说明是数组的指针；如p [5]*int，*与int结合，说明这个数组都是int类型的指针，是指针数组
	fmt.Println("---------数组指针----------")
	i := [...]int{9: 1}
	var p *[10]int = &i
	fmt.Println(p)
	fmt.Println(*p)

	x, y := 1, 2
	j := [...]*int{&x, &y}
	fmt.Println(j)

	//判断数组是否相等，数组之间可以用==或！=比较，但是不能用<和>
	k := [2]int{1, 2}
	l := [2]int{1, 2}
	fmt.Println("k == 1:", k == l)

	//指向数组的指针
	//可以用 new 来创建一个数组，该方法返回一个指向该数组的指针
	m := new([5]int)
	fmt.Println("m: ", m)

	//对某个元素进行操作
	aaa := [10]int{}
	aaa[1] = 2
	fmt.Println("aaa: ", aaa)
	bbb := new([10]int)
	bbb[1] = 2
	fmt.Println("bbb[1]: ", bbb[1])
	fmt.Println("bbb: ", bbb)
	fmt.Printf("&aaa: %v\n", &aaa)

	ccc := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println("ccc: ", ccc)

	ddd := [...][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println("ddd: ", ddd)

	fff := [...][3]int{
		{1, 1},
		{0: 2, 2: 2}}
	fmt.Println("fff: ", fff)
}

func change_Array(carray [5]int) {
	carray[0] = 100
}

func change_Slice(cslice []int) {
	cslice[0] = 100
}

//获取变量类型
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
