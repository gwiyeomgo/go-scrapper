package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var errRequestFailed = errors.New("Request failed")

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

/*func main() {
//url 을 동시에 check 하고 싶다
//하나씩 순서대로 확인
urls := []string{
	"https://www.google.com",
	"https://www.amazon.com",
	"https://www.airbnb.com",
	"https://www.facebook.com",
}
//var results  map[string]string
//초기화 되지 않은 results 값에는 값을 넣을 수 없다
//ex) results["hello"] = "hello"
//empty map에 값을 넣고 싶다면
//(1)var results= map[string]string{}
//(2)
var results = make(map[string]string)
for _,url := range  urls {
	result :="OK"
	err := hitUrl(url)
	if err != nil {
		result="failed"
	}
	results[url]= result
}
for url,result := range results {
	fmt.Println(url,result)
}

}
*/
//goroutines 기본적으로 다른 함수와 동시에 실행시키는 함수
//go를 통해서 원하는건 좀 더 빠른 방식의 프로그래밍이다

/*func main()  {
	//main 함수가 끝이 나면 모든 goroutine 도 끝이 나는 거다
	//메인 함수는 goroutines 를 기다려주지 않는다
	//우리가 goroutine 으로부터 main function 까지 어떻게 커뮤니케이션 해야 할지 모른다
    // main 함수에게 결과를 보내주는 방법은 무엇일까?
	//Channels
	//Channel은 goroutine이랑 메인함수 사이에 정보를 전달하기 위한 방법이다
	//또는 goroutine에서 다른 goroutine으로 커뮤니케이션하는것도 가능하다
	//Channel은 파이프 같은거다
	// 너가 이 파이프를 통해서 메세지를 보낼수도,받을 수 도 있다
   // goroutine이 main함수로 메세지를 보내고 싶을때 channel을 사용한다
	go sexyCount("mini")
	go sexyCount("tina")
	time.Sleep(time.Second *10)
	//channel 만들기
	//c := make(chan bool)
	c := make(chan string)
	people := [3]string{"mini","tina","ann"}
	for _,person := range people {
		//체널을 함수로 보내는 방법
		go isSexy(person,c)
	}
	//main은 channel이 보낸 메세지를 어떻게 알지?
	//아래처럼 쓴다면(채널로부터 뭔가를 받을때) main 함수가 어떤 답이 올때까지 기다린다
	//blocking operation =메세지를 받을떄까지 기다린다
	//이 작업이 끝날때까지 멈춘다는 뜻
	result := <-c //채널로부터 메세지를 가져온다
	fmt.Println(result)
	//우리는 2개 메세지를 보냈으니 2개 메세지를 받을 수 있다
	//fmt.Println("Waiting for messages")
	//fmt.Println("Received this message:",<-c)//true
	//fmt.Println("Received this message:",<-c)//true
	//fmt.Println(<-c)//fatal error: all goroutines are asleep - deadlock!
    //main이 메세지를 기다리고 있는데 실제로 만든 goroutines는 끝나버렸다

	//people 숫자가 늘어난다면 loop 사용해서 goroutine의 수만큼 메세지를 받을 수 있다.
	//3개의 메세지 리시버를 만듬
	for i:=0;i<len(people);i++ {
		fmt.Println(<-c)
	}
}*/

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	//프로그램이 끝나는 시간은
	//가장 체크가 오래걸리는 url 하나에 걸리는 시간이다
	urls := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.airbnb.com",
		"https://www.facebook.com",
	}
	for _, url := range urls {
		go hitUrl(url, c)
	}
	for i := 0; i < len(urls); i++ {
		//fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

//만약 channel로 연결된 함수에
/*func hitUrl(url string ,c chan result)  {
	fmt.Println("checking:",url)
	//아래와 같이 코드를 쓴다면
	//해당 함수는 받을 수만 있고 메세지를 보낼 수 없다고 설정하는 것이다(send only)
	fmt.Println(<-c)
}*/
///*func hitUrl(url string ,c chan<- result)  {
//`c chan<- result`써도 같은 의미 send only
func hitUrl(url string, c chan<- requestResult) {
	fmt.Println("checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}

// c chan bool
// (채널이름) (채널타입) (그 채널을 통해서 어떤 타입의 데이터를 주고 받을지 )
//func isSexy(person string,c chan bool) {
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	//* how does goroutine send a message using that channel ?
	//c <- true
	c <- person + " is sexy"
}
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

/*func hitUrl(url string) error {
	fmt.Println("checking:",url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}*/
