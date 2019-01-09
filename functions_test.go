package main

import (
	"fmt"
	libf "libfunc"
	"reflect"
	"testing"
)

/*
	函数
*/
func Test_Functions(t *testing.T) {
	res := plus(3, 5)
	fmt.Println("3 + 5 =", res)

	//调用多返回值函数
	//多赋值操作
	a, b := vals()
	fmt.Println("a b = ", a, b)

	_, c := vals()
	fmt.Println("c =", c)

	//调用变参函数
	sum(1, 2)
	sum(1, 2, 3)

	//如果 slice 已经有了多个值，想把它们作为变参使用，要这样调用 func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	f_fun(2, "Go", 8, "language", 'a', false, 'A', 3.14)

	fmt.Println(reflect.TypeOf('a').String())

}

//fmt.Print()函数可以传递不同类型的参数，go语言规定，如果希望传递任意类型的变参，变参类型应该制定为空接口类型：interface{}.
func f_fun(args ...interface{}) {
	var num = make([]int, 0, 6)
	var str = make([]string, 0, 6)
	var ch = make([]int32, 0, 6)
	var other = make([]interface{}, 0, 6)

	/*类型switch语句与一般形式有两点差别。第一点，紧随case关键字的不是表达式，而是类型说明符。类型说明符由若干个类型字面量组成，
	 *且多个类型字面量之间由英文逗号分隔。第二点，它的switch表达式是非常特殊的。这种特殊的表达式也起到了类型断言的作用，但其表现形式
	 *很特殊，如：v.(type) , 其中v必须代表一个接口类型的值。该类表达式只能出现在类型switch语句中，且只能充当switch表达式。
	 */
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			num = append(num, v)
		case string:
			str = append(str, v)
		case int32:
			ch = append(ch, v)
		default:
			other = append(other, v)
		}

	}

	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", num)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", str)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", ch)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", other)
}

func plus(a int, b int) int {
	return a + b
}

//多返回值
func vals() (int, int) {
	return 3, 5
}

//变参函数: 函数中形式参数的数目通常是确定的，在调用的时候要依次传入与形式参数对应的所有实际参数，但是在某些函数的参数个数可以根据实际需要来确定，这就是变参函数。
//        Go语言支持不定长变参，但是要注意不定长参数只能作为函数的最后一个参数，不能放在其他参数的前面
//不定长变参的实质就是一个切片，可以使用range遍历
func sum(nums ...int) {
	fmt.Print(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
