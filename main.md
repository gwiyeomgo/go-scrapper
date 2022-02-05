package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
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

/*func main() {
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
}*/

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

var baseUrl string = "https://kr.indeed.com/jobs?q=go&limit=50"

type extractedJob struct {
	id       string
	title    string
	location string
	company  string
	summary  string
}

//https://github.com/PuerkitoBio/goquery 를 사용해서 html 의 내부를 들여다볼 수 있게 해준다
func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
		extractedJobs := getPage(i)
		//2개 배열을 합친다
		//배열안에 넣은거 아님
		jobs = append(jobs, extractedJobs...)
	}
	writeJob(jobs)
}

//https://pkg.go.dev/encoding/csv
// []extractedJob 을 csv 파일로 변환
func writeJob(jobs []extractedJob) {
	//파일을 생성한다
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	//w.Flush() 함수가 끝나는 시점에 파일에 데이터를 입력하는 함수
	defer w.Flush()

	headers := []string{
		"Link",
		"Title",
		"Location",
		"Company",
		"Summary",
	}
	//headers 내용으로 파일이 만들어진다.
	err = w.Write(headers)
	checkErr(err)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.company, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

//특정 페이지 정보를 가져온다
func getPage(page int) []extractedJob {
	var jobs []extractedJob
	//	pageUrl := baseUrl +"&stat=" +page * 50
	//string 과 int 를 같이 쓰려면 정수를 string으로 바꾸는 strconv.Itoa 를 쓴다
	pageUrl := baseUrl + "&stat=" + strconv.Itoa(page*50)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	checkErr(err)
	//20220205 기준
	//a 태그의 data-jk 값을 가져온다.
	searchCards := doc.Find("#mosaic-provider-jobcards")
	searchCards.Find(".tapItem").Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		//추출된 struct 를 배열에 요소로 추가하고 있다.
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	company := cleanString(card.Find(".companyName").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	return extractedJob{id: id, title: title, location: location, company: company, summary: summary}
}

//전체 웹사이트 페이지를 가져온다
func getPages() int {
	pages := 0
	res, err := http.Get(baseUrl) //resp *Response
	checkErr(err)
	checkCode(res)

	//res.Body 는 기본적으로 byte인데,입력과 출력이야 (io)
	//그래서 res.Body를 닫아야 한다
	//getPages 함수가 끝났을때
	//아래처럼 쓴다면 메모리가 새어나가는 걸 막을 수 있다
	defer res.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//http.Get 의 반환 타입은 *Response 이다
//따라서 checkCode의 매개변수의 타입을 아래와 같이 지정한다.
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
func cleanString(str string) string {
	//golang strings
	//공백을지우고 모든 단어를 분리된 글자로 만드는 것
	//Fields 가 string 으로 된 배열을 반환
	//모든 공백이 제거됨 글자만 배열에
	//또한 Join을 사용한다
	//"hello" "a" "b" => hello a v
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
