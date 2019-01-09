package main

import (
	"fmt"
	"testing"
)

/*
  你好，世界
*/
func Test_Hello(t *testing.T) {
	h := "hello, world"
	fmt.Printf("%v", h)
}
