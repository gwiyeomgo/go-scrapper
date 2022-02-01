package main

import (
	"fmt"
	"github.com/gwiyeomgo/go-scrapper/accounts"
	"log"
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
	/*fmt.Println("Hello world")
	something.SayHello() //다른 패키지에서 export 된 func

	totalLength, upperName := lenAndUpper("test")
	fmt.Println(totalLength, upperName)

	repeatMe("A", "B", "C", "D", "E")*/

	// go 에서 constructor 만드는 방법
	account := accounts.NewAccount("gwiyeomgo")
	//fmt.Println(account)
	//&{gwiyeomgo 0}
	//복사본을 return 하는 것이 아닌
	//실제 메모리 주소(address)가 출력된다. 복사본이 아니라 object 를 return 시켰다
	account.Deposit(10) //입금하다
	//fmt.Println(account.Balance()) //10
	//account.Withdraw(20) //인출하다
	//fmt.Println(account.Balance()) //-10
	//error 를 다루기 위해서
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err) // println을 호출하고 프로그램 종료 시킴
	}
}
