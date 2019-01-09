package main

import (
	"fmt"
	"testing"
)

var I int

/*
* init函数在每个package是可选的，可有可无，甚至可以有多个(但是强烈建议一个package中有一个init函数)，
* init函数在你导入该package时程序会自动调用init函数，所以init函数不用我们手动调用,另外它只会被
* 调用一次，因为当一个package被多次引用时，它只会被导入一次
 */
//func init() {
//	I = 0
//	fmt.Println("Call mypackage init1")
//}

//func init() {
//	I = 1
//	fmt.Println("Call mypackage init2")
//}

/*
方法：在结构体类型中定义方法
	！Go中虽没有class，但依旧有method
	！通过显示说明receiver来实现与某个类型的结合
	！只能为同一个包中的类型定义方法
	！receiver可以是类型的值或者指针
	！不存在方法重载
	！可以使用值或指针来调用方法，编译器会自动完成转换
	！从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第一个参数(Method Value vs. Method Expression)
	！如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
	！类型别名不会拥有底层类型所附带的方法
	！方法可以调用结构中的非公开字段
*/

type rect struct {
	width, height int
}

type A struct {
	Name string

	//首字母小写代表私有字段
	mm string
}

type B struct {
	Name string
}

type myint int

func Test_Methords(t *testing.T) {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	//结果所示：值类型不使用指针，在这个方法结束之后，值不会被修改
	/*
	*	A
	*	AA
	*	B
	 */
	a := A{}
	a.Print()
	fmt.Println(a.Name)
	b := B{}
	fmt.Printf("&b: %p\n", &b)
	b.Print()
	fmt.Println(b.Name)

	var i myint = 2
	i.myDouble()
	fmt.Println("i=", i)
}

//编译器根据接收者的类型，来判断它是属于哪个方法
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

//编译器根据接收者的类型，来判断它是属于哪个方法
//加上*代表指针传递，如果想在方法中修改对象的值只能用pointer receiver，对象较大时避免拷贝也要用pointer receiver
func (a *A) Print() {
	//取一个变量a，a就是接收者，它的接收者的类型就是structA,Print就是方法的名称，参数在Print()的括号中定义
	//receiver就是这个函数的第一个接收者，而且是强制规定的，这个时候就变成了一个方法
	a.Name = "AA"
	fmt.Println("A")

	//方法访问权限:可以调用结构中的私有字段
	a.mm = "123"
	fmt.Println("a.mm: ", a.mm)
}

//这里的b并不是以指针传递
func (b B) Print() {
	fmt.Printf("&bt: %p\n", &b)
	b.Name = "BB"
	fmt.Println("B")
}

func (p *myint) myDouble() int {
	*p = *p * 2
	return 0
}
