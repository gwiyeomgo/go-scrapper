package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var baseUrl string = "https://kr.indeed.com/jobs?q=go&limit=50"

type extractedJob struct {
	id       string
	title    string
	location string
	company  string
	summary  string
}

func main() {
	var jobs []extractedJob
	mainC := make(chan []extractedJob)
	//이 채널로 일자리 정보 여러개가 전달된다
	//(1) 총 페이지 수를 가져와
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		//goroutine 키워드 입력, channel을 전달
		//총 5개 channel
		go getPage(i, mainC)
	}

	//gorutine을 기다려야 함
	//totalPage만큼
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-mainC
		jobs = append(jobs, extractedJobs...)

	}

	writeJob(jobs)
}

//getPage가 값을 리턴 x
//main에서 만든 채널로 갑을 전송하도록 변경
func getPages() int {
	pages := 0
	res, err := http.Get(baseUrl)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
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

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageUrl := baseUrl + "&stat=" + strconv.Itoa(page*50)
	fmt.Println(pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	checkErr(err)
	searchCards := doc.Find("#mosaic-provider-jobcards").Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	//go extractJob gorutine 생성했고 값을 전달받는다
	// <-c searchCards 안에 card 숫자만큼 생성해야 함
	//전달받을 메세지의 숫자는 카드의 갯수와 같다
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	company := cleanString(card.Find(".companyName").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	c <- extractedJob{id: id, title: title, location: location, company: company, summary: summary}
}

func writeJob(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{
		"Link",
		"Title",
		"Location",
		"Company",
		"Summary",
	}
	err = w.Write(headers)
	checkErr(err)

	//file을 동시에 작성하도록 수정
	/*for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.company, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}*/
	writeC := make(chan []string)
	/*	wireFunction := func(job extractedJob) {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.company, job.summary}
		writeC <- jobSlice
	}*/
	for _, job := range jobs {
		go func(job extractedJob) {
			jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.company, job.summary}
			writeC <- jobSlice
		}(job)
		//go wireFunction(job)
	}

	for i := 0; i < len(jobs); i++ {
		jobSlice := <-writeC
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)

	}

}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
