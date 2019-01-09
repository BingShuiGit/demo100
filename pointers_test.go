package main

import (
	"fmt"
	"testing"
)

const MAX int = 3

/*
	指针：允许在程序中通过引用传递值或者数据结构。通过两个函数：zeroval 和 zeroptr 来比较指针和值类型的不同。
*/
func Test_Pointers(t *testing.T) {
	ii := 1
	fmt.Println("&ii: ", &ii)
	fmt.Println("initial:", ii)

	//i的值并没有发生改变，还是1（值传递）<-------------->（与map中结构体方法参数传递对比）
	zeroval(ii)
	fmt.Println("zeroval:", ii)

	//i的值发生了改变，改为0（改变这个指针引用的真实地址的值）
	zeroptr(&ii)
	fmt.Println("zeroptr:", ii)

	fmt.Println("pointer:", &ii)

	//三个整数将存储在指针数组中
	aaa := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		//整数地址赋值给指针数组
		ptr[i] = &aaa[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("aaa[%d] = %d\n", i, *ptr[i])
	}

	// 调用函数用于交换值
	// &a 指向 a 变量的地址
	// &b 指向 b 变量的地址
	var pt_a int = 100
	var pt_b int = 200

	swap(&pt_a, &pt_b)
	fmt.Printf("交换后 pt_a 的值 : %d\n", pt_a)
	fmt.Printf("交换后 pt_b 的值 : %d\n", pt_b)
}

func zeroval(ival int) {
	fmt.Println("&ival: ", &ival)
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

//交换两个数的值（指针作为函数参数）
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
