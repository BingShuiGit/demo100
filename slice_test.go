package main

import (
	"fmt"
	"testing"

	libf "libfunc"
)

/*
	切片：比数组好用的类型
*/
func Test_Slice(t *testing.T) {
	//切片声明1
	var ss []string
	fmt.Printf("length:%v \taddr:%p \tisnil:%v\n", len(ss), ss, ss == nil)

	//切片尾部追加元素append elemnt
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("S%d", i))
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\tssCap:%v\tss:%v\tisnil:%v\n", len(ss), " ssCap: ", ss, ss == nil)
	print("after append\t", ss)

	//删除切片元素remove element at index , append(切片， 元素)元素是可变长参数加上...表示
	//最后的“...”省略号表示接收变长的参数为slice
	index := 5
	ss = append(ss[:index], ss[index+1:]...)
	print("after delete", ss)

	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "inserted")
	ss = append(ss, rear...)
	print("after insert", ss)

	//切片声明2
	s := make([]string, 3)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "s: ", s, "len: ", len(s), "cap: ", cap(s))

	//切片声明3
	r := []string{"g", "h", "i"}
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Dcl: ", r, "rLen: ", len(r), "rCap: ", cap(r))

	//cap(s),len(s)
	s2 := s[:]
	s3 := s
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "s2: ", len(s2), cap(s2), s2, "\ns3: ", len(s3), cap(s3), s3, "\n")

	//切片赋值
	s[0], s[1], s[2] = "a", "b", "c"
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Set: ", s)
	fmt.Println("Get: ", s[2])

	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Len: ", len(s), "Cap: ", cap(s))

	//混合类型切片
	in := make([]interface{}, 0)
	in = append(in, 2, "Go", 8, "language", 'a', false, "A", 3.14)
	fmt.Println("in: ", in)
	fmt.Println("----------------------------------------------")

	//切片的追加
	s = append(s, "d")
	s = append(s, "e", "f")
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Append: ", s, "Len: ", len(s), "Cap: ", cap(s))

	//切片的复制,copy的使用
	c := make([]string, len(s))
	copy(c, s)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Copy: ", c, "CLen: ", len(c), "CCap: ", cap(c))

	l := s[2:5]
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Sli1: ", l, "lLen1: ", len(l), "lCap1: ", cap(l))

	//在使用copy复制切片之前，要保证目标切片有足够的大小，注意是大小，而不是容量
	var sa = make([]string, 0)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))

	}

	//此时da大小为0，容量为10
	var da = make([]string, 0, 10)

	//这里的cc作用是什么？
	var cc = 0
	cc = copy(da, sa)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "cc: ", cc)
	fmt.Printf("copy to da(len=%d)\t%v\n", len(da), da)

	//此时da大小是5，复制后[0,1,2,3,4]
	da = make([]string, 5)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

	//此时da大小是5，复制后[0,1,2,3,4,5,6,7,8,9]
	da = make([]string, 10)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

	//当给出初始大小后，得到的实际上是一个含有这个size数量切片类型的空元素, 大小为10，容量默认和大小一样，为10
	var sc = make([]string, 10)
	//追加后大小是11，前10个位空元素
	sc = append(sc, "last")
	print("sc after append: ", sc)

	//使用append函数而不想区分是否有空的元素，大小为0，容量为10
	var sd = make([]string, 0, 10)
	sd = append(sd, "last")
	print("sd after append", sd)

	//以下 取s其中一段，子切片做append， 添加1个，多个，超过s容量
	l = append(l, "f")
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "\tSli11: ", l, "\n\t\tlLen11: ", len(l), "lCap11: ", cap(l))
	l = append(l, "g", "gg")
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "\tSli12: ", l, "\n\t\tlLen12: ", len(l), "lCap12: ", cap(l))
	l = append(l, "g", "h", "i", "j", "k")
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "\tSli13: ", l, "\n\t\tlLen13: ", len(l), "lCap13: ", cap(l))

	l = s[:5]
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Sli2: ", l)

	l = s[2:]
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "Sli3: ", l)

	//当用append追加元素到切片时，如果容量不够，go就会创建一个新的切片变量，看下面程序的执行结果：addr一直在改变
	var sb []string
	fmt.Printf("addr:%p \t\t\tlen:%v content:%v\n", sb, len(sb), sb)
	for i := 0; i < 10; i++ {
		sb = append(sb, fmt.Sprintf("%v", i))
		fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sb, len(sb), sb)
	}
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n\n", sb, len(sb), sb)

	//如果在make初始化切片时给出了足够的容量，append操作不会创建新的切片：与上面对比addr
	var sbb = make([]string, 0, 10)
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sbb, len(sbb), sbb)
	for i := 0; i < 10; i++ {
		sbb = append(sbb, fmt.Sprintf("%v", i))
		fmt.Printf("addr:%p \t\tlen:%v content:%v\n", sbb, len(sbb), sbb)
	}
	fmt.Printf("addr:%p \t\tlen:%v content:%v\n\n", sbb, len(sbb), sbb)

	//切片指针:如果不能准确预估切片的大小，又不想改变变量,这时候使用指针。
	//本质上是：append操作亦然会在需要的时候构造新的切片，不过是将地址都保存到了posa中，因此我们通过该指针始终可以访问到真正的数据。
	var osa = make([]string, 0)
	posa := &osa
	for i := 0; i < 10; i++ {
		*posa = append(*posa, fmt.Sprintf("%v", i))
		fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", osa, posa, posa)
	}
	fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n\n", osa, posa, posa)

	//调用参数为切片的函数
	sli4 := []int{0, 1, 2, 3}
	fmt.Printf("sli4: %v sli4addr: %p \n", sli4, &sli4)

	//和map一样，值会被改掉
	ret := changeSlice1(sli4)
	fmt.Println("ret: ", ret)
	fmt.Printf("sli4: %v &sli4: %p retaddr: %p \n", sli4, &sli4, ret)

	//调用参数为切片指针的函数
	sli5 := []int{0, 1}
	fmt.Printf("sli5 %v %p \n", sli5, &sli5)

	changeSlice2(&sli5)
	fmt.Printf("sli5 %v %p \n", sli5, &sli5)

	sli5[1] = -1111
	fmt.Printf("sli5 %v %p \n", sli5, &sli5)

	//二维切片
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	for i := 0; i < 3; i++ {
		//第i行有i+1列
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		fmt.Println(twoD[i])
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)

	twoD2 := make([][]int, 3)
	fmt.Println(twoD2)
	for i := 0; i < 3; i++ {
		twoD2[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD2)
}

//切片做函数参数
func print(msg string, ss []string) {
	fmt.Printf("[ %20s ]\t:\tlength:%v\taddr:%p\tisnil:%v\tcontent:%v", msg, len(ss), ss, ss == nil, ss)
	fmt.Println()
}

//当切片 sli4 当做参数 s 传进来时，和 main 函数中的sli4不是同一个，由func1和sli4addr可以看出
func changeSlice1(s []int) []int {
	fmt.Printf("func1: %p \n", &s)
	s[1] = 111
	fmt.Printf("func2: %p \n", &s)
	return s
}

//切片指针作为函数参数：当切片地址 &sli5 当做参数 s 传进来时，和 main 函数中的sli5是同一个；
//传递给函数的是sli5的指针，函数内对 s 的操作本质上都是对sli5的操作。并且也可以从函数内打出的 s 地址看到，至始至终就只有一个切片
func changeSlice2(s *[]int) {
	fmt.Printf("func s %v %p \n", *s, s)
	(*s)[0] = -1
	*s = append(*s, 3)
	(*s)[1] = 1111
}

type Struct1Stu struct {
	Int1 int
	Str1 string
	F1   float64
	// +map，因为是键值对，会像引用传递一样改变所有的值
	Map1      map[int]string
	IntSclice []int
	IntArry   [5]int
}

func (this *Struct1Stu) Clone() (cln Struct1Stu) {
	cln = *this
	cln.IntSclice = make([]int, len(this.IntSclice))
	copy(cln.IntSclice, this.IntSclice)

	return
}

type Struct2Stu struct {
	Int2       int
	Str2       string
	F2         float64
	Map2       map[int]string
	IntSclice2 []int
	IntArry2   [5]int

	Struct2 Struct1Stu
}

func (this *Struct2Stu) Clone() (cln Struct2Stu) {

	cln = *this
	cln.IntSclice2 = make([]int, len(this.IntSclice2))
	//显示新建了一个切片地址来存放初始切片
	fmt.Printf("cln.IntSclice2 addr: %p len: %d", &cln.IntSclice2, len(cln.IntSclice2))
	fmt.Printf("this.IntSclice2 addr: %p len: %d", &this.IntSclice2, len(this.IntSclice2))
	copy(cln.IntSclice2, this.IntSclice2)

	cln.Struct2 = this.Struct2.Clone()

	return
}

//混合赋值修改，不适用Clone（），修改后struct22的slice类型数据改变引起struct21的slice类型数据改变
func Test_SliceStruct(t *testing.T) {

	var struct21, struct22, struct23 Struct2Stu
	var struct1 Struct1Stu = Struct1Stu{
		Int1:      11,
		Str1:      "str1",
		F1:        11.1,
		Map1:      map[int]string{1: "a", 2: "b"},
		IntSclice: []int{11, 12},
		IntArry:   [5]int{101, 102},
	}

	struct21.Int2 = 2
	struct21.Str2 = "ddd"
	struct21.F2 = 2.23
	struct21.Map2 = map[int]string{1: "A", 2: "B"}
	struct21.IntSclice2 = []int{21, 22}
	struct21.IntArry2[1] = 22

	struct21.Struct2 = struct1

	struct22 = struct21
	struct23 = struct21.Clone()

	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "struct21=%v", libf.StructToJsonStringOne(struct21))
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "struct22=%v", libf.StructToJsonStringOne(struct22))
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "struct23=%v", libf.StructToJsonStringOne(struct23))

	struct21.Int2 = 555555555
	struct21.IntArry2[1] = 2222

	struct21.Map2[1] = "C"
	struct21.Map2[2] = "D"

	struct21.Struct2.Map1[1] = "c"
	struct21.Struct2.Map1[2] = "d"

	struct21.IntSclice2[0] = 211111
	struct21.IntSclice2[1] = 221111
	struct21.IntSclice2 = append(struct21.IntSclice2, 31111)

	struct21.Struct2.Map1[1] = "aaaaaa"
	struct21.Struct2.Map1[2] = "bbbbbb"
	struct21.Struct2.Map1[3] = "cccccc"

	struct21.Struct2.IntSclice[0] = 11555555
	struct21.Struct2.IntSclice[1] = 12555555
	struct21.Struct2.IntSclice = append(struct21.Struct2.IntSclice, 1212121)

	struct21.Struct2.IntArry[0] = 1011111
	struct21.Struct2.IntArry[1] = 1021111

	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "struct21=%v", libf.StructToJsonStringOne(struct21))
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "struct22=%v", libf.StructToJsonStringOne(struct22))
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "struct23=%v", libf.StructToJsonStringOne(struct23))

	return
}
