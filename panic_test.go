package main

import (
	//	libf "libfunc"
	"os"
	"testing"
)

func Test_Panic(t *testing.T) {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
