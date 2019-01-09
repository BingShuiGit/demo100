package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func Test_Http(t *testing.T) {
	//注册URI路径与相应的处理函数
	http.HandleFunc("/hello", sayHelloHttp)
	// 监听9090端口，就跟javaweb中tomcat用的8080差不多一个意思吧
	er := http.ListenAndServe(":9091", nil)
	if er != nil {
		log.Fatal("ListenAndServe: ", er)
	}
}

func sayHelloHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}
