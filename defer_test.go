package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_Defer(t *testing.T) {
	f_Defer := createFile("/tmp/defer.txt")
	defer closeFile(f_Defer)
	writeFile(f_Defer)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
