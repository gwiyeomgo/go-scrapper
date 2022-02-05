package scrapper

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

type extractedJob struct {
	id       string
	title    string
	location string
	company  string
	summary  string
}

func Scrapper(term string) {
	var baseUrl string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	mainC := make(chan []extractedJob)
	totalPages := getPages(baseUrl)

	for i := 0; i < totalPages; i++ {
		//goroutine 키워드 입력, channel을 전달
		//총 5개 channel
		go getPage(i, baseUrl, mainC)
	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-mainC
		jobs = append(jobs, extractedJobs...)

	}

	writeJob(jobs)
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
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

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageUrl := url + "&stat=" + strconv.Itoa(page*50)
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

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".jobTitle>span").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	company := CleanString(card.Find(".companyName").Text())
	summary := CleanString(card.Find(".job-snippet").Text())
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

	writeC := make(chan []string)
	for _, job := range jobs {
		go func(job extractedJob) {
			jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.company, job.summary}
			writeC <- jobSlice
		}(job)
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

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
