package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

//Go 提供内置的 JSON 编解码支持，包括内置或者自定义类型与 JSON 数据之间的转化
//Json可以用有效可读的方式表示基础数据类型和数组、slice、结构体、map等聚合数据类型
func Test_Json(t *testing.T) {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("mapB: ", string(mapB))

	//JSON 包可以自动的编码你的自定义类型。编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
	//将一个Go语言中结构体转为JSON的过程叫编组(marshaling),编组通过调用json.Marshal函数完成(data, err := json.Marshal(结构体实例))
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	res2D := Response2{
		Page:   2,
		Fruits: []string{"apple", "peach", "pear"}}

	//可以给结构字段声明标签来自定义编码的 JSON 数据键名称(15/16行)
	//结果中键变成了小写fruit是因为16行的`json:"fruits"`
	//json.MarshalIndent函数将产生整齐缩进的输出。该函数有两个额外的字符串参数用于表示每一行输出的前缀和 *每一个层级* 的缩进
	res2B, _ := json.MarshalIndent(res2D, "", "	")
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//map[string]interface{} 将保存一个 string 为键，值为任意值的map
	var dat map[string]interface{}

	//将JSON转换为Go数据类型:func Unmarshal(data []byte, v interface{}) error; 此函数将data表示的JSON转换为v
	//解码和相关的错误检查
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat:", dat)
	//为了使用解码 map 中的值，我们需要将他们进行适当的类型转换,将 num 的值转换成 float64类型
	num := dat["num"].(float64)
	fmt.Println(num)
	//访问嵌套的值需要一系列的转化
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	//编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，通过json.Unmarshal函数完成
	//JSON中的字段在Go目标类型中不存在， json.Unmarshal()函数在解码过程中会丢弃该字段
	json.Unmarshal([]byte(str), &res)
	fmt.Println("res: ", res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	fmt.Println("d:", d)
}
