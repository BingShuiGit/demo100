package main

import (
	"fmt"
	sr "strings"
	"testing"
)

//给 fmt.Println 一个短名字的别名
var ptn = fmt.Println

func Test_String(t *testing.T) {
	ptn("Contains:  ", sr.Contains("test", "es"))
	ptn("Count:     ", sr.Count("test", "t"))
	ptn("HasPrefix: ", sr.HasPrefix("test", "te"))
	ptn("HasSuffix: ", sr.HasSuffix("test", "st"))
	ptn("Index:     ", sr.Index("test", "e"))
	ptn("Join:      ", sr.Join([]string{"a", "b", "c"}, "-"))
	ptn("Repeat:    ", sr.Repeat("a", 5))
	//将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	ptn("Replace:   ", sr.Replace("foo", "o", "0", -1))
	//将出现的第一个o变为0
	ptn("Replace:   ", sr.Replace("foo", "o", "0", 1))
	ptn("Split:     ", sr.Split("a-b-c-d-e", "-"))
	ptn("ToLower:   ", sr.ToLower("TEST"))
	ptn("ToUpper:   ", sr.ToUpper("test"))
	ptn()
	//获取字符串长度和通过索引获取一个字符
	ptn("Len: ", len("hello"))
	ptn("Char:", "hello"[1])

}
