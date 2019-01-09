package main

import (
	"errors"
	"fmt"
	libf "libfunc"
	"math"
	"testing"
)

//定义一个几何体的基本接口
type geometry interface {
	area() float64
	perim() float64
}

type getGirl interface {
	getName() string
	getAge() int
}

//接口
type IFile interface {
	Read()
	Write()
}

type IReader interface {
	Read()
}

type File struct {
}

type girl struct {
	name string
	age  int
}

type rectt struct {
	width, height float64
}

type circle struct {
	radius float64
}

type item struct {
	Name string
}

func (i item) String() string {
	return fmt.Sprintf("item name: %v", i.Name)
}

type person_interfaces struct {
	Name string
	Sex  string
}

func (p person_interfaces) String() string {
	return fmt.Sprintf("person name: %v sex: %v", p.Name, p.Sex)
}

/*
接口: 接口类型是由一组方法定义的集合，任何类型的方法集 中只要拥有与之对应的全部方法，就表示它实现了该接口。
*/
func Test_Interfaces(t *testing.T) {
	r := rectt{width: 3, height: 5}
	c := circle{radius: 5}

	var ir geometry = &r
	fmt.Println("ir: ", ir)

	measure(ir)
	measure(c)

	var xiao girl
	xiao.name = "xiaosan"
	xiao.age = 38
	getGirlInfo(xiao)

	san := &girl{name: "夫子庙小三", age: 38}
	getGirlInfo(san)

	f := new(File)

	// ok 因为FIle实现了IFile中的所有方法  ,如果没有赋值f，会报panic错误，因为调用一个空接口值上的任意方法都会产生panic
	var f11 IFile = f
	f11.Read()

	//接口赋值
	//一个接口的值，接口值，由两个部分组成，一个具体的类型和该类型的值，它们被称为接口的动态类型和动态值
	// ok 因为IFile中包含IReader中所有方法
	var f22 IReader = f11

	// error 因为IReader并不能满足IFile（少一个方法）
	// var f3 IFile = f2

	// ok 因为File实现了IReader中所有方法
	var f33 IReader = new(File)

	// error 因为IReader并不能满足IFile 同上..如何解决呢？ 要用接口查询
	// var f44 IFile = f33

	// 接口查询
	// 这个if语句检查接口指向的对象实例 是否 实现了IFile接口
	// 如果实现了则执行特定的代码。
	// 注意：这里强调的是对象实例，也就是new(File)
	// File包含IFile里所有的方法
	// 所以ok = true
	if f55, ok := f33.(IFile); ok {
		prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "f55: ", f55)
	}

	// 询问接口它指向的对象 是否 是某个类型
	// 这个if语句判断接口指向的对象实例是否是*File类型， 是的话返回f33的动态值
	// 依然ok
	if f66, ok := f33.(*File); ok {
		prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "f66: ", f66)
	}

	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "f11: ", f11, "f22: ", f22, "f33: ", f33)

	p1 := Parse("Apple").(*item)
	fmt.Println(p1)
	p2 := Parse([]string{"zhangsan", "man"}).(*person_interfaces)
	fmt.Println(p2)

}

//让 rect 和 circle 实现这个接口
func (r rectt) area() float64 {
	return r.width * r.height
}

func (r rectt) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//girl结构体实现
//(1) 类型T的对象赋值给接口变量时，要求类型T的所有方法的接收器类型必须是T
//(2) 类型*T的对象赋值给接口变量时，类型T的方法接收器类型可以是T或者*T,原因是：
//		Go语言可以根据函数：func (g girl) getName() string自动生成一个新的getName()方法：func (g *girl) getName() string
//		但是不能根据函数：func (g *girl) getName() string自动生成方法：func (g girl) getName() string
func (g girl) getName() string {
	return g.name
}

func (g girl) getAge() int {
	return g.age
}

func (f *File) Read() {
	fmt.Println("Read")
}

func (f *File) Write() {
	fmt.Println("Write")
}

//接口作为参数，实现该接口的结构体可以调用接口中的方法
func measure(g geometry) {
	//这里是怎么打印出结构体的 g不是接口么？ 到measure的时候，g自动获取结构体
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func getGirlInfo(gg getGirl) {
	fmt.Println("gg: ", gg)
	fmt.Println("girlName is ", gg.getName())
	fmt.Println("girlAge is ", gg.getAge())
}

//interface{}作为函数返回值
func Parse(i interface{}) interface{} {
	switch i.(type) {
	case string:
		return &item{
			Name: i.(string),
		}
	case []string:
		data := i.([]string)
		length := len(data)
		if length == 2 {
			return &person_interfaces{
				Name: data[0],
				Sex:  data[1],
			}
		} else {
			return nil
		}
	default:
		panic(errors.New("type match miss"))
	}
	return nil
}
