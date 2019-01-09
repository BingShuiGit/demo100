package main

import (
	"fmt"
	"testing"
)

/*
  值
*/
func Test_Value(t *testing.T) {
	fmt.Println("go" + "lang" + "\n")

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println("7/3=", 7/3)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
