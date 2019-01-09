package main

import (
	"fmt"
	"math"
	"testing"
)

/*
  常量
*/
const s string = "constant"

func Test_Constant(t *testing.T) {
	fmt.Println(s)

	//类型的转换
	const n = 500000000
	fmt.Println(float64(n))

	//3e20和3e+20一样，指3 * 10的20次方
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	//进去math类的Sin方法可以看到Sin()需要float64类型的参数
	fmt.Println(math.Sin(float64(n)))
}
