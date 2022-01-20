package main

import (
	"fmt"
	"github.com/gwiyeomgo/go-scrapper/something"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	//go는 다양한 package 를 갖고있다.(표준 라이브러리 확인)
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	//여러개의 arguments 전달
	fmt.Println(words)
}
func main() {
	fmt.Println("Hello world")
	something.SayHello() //다른 패키지에서 export 된 func

	totalLength, upperName := lenAndUpper("test")
	fmt.Println(totalLength, upperName)

	repeatMe("A", "B", "C", "D", "E")
}
