package main

import (
	"github.com/gwiyeomgo/go-scrapper/scrapper"
	"github.com/labstack/echo"
	"os"
	"strings"
)

const fileName string = "jobs.csv"

//https://echo.labstack.com/guide/
//echo 설치
//go get -u github.com/labstack/echo
func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	e.Logger.Fatal(e.Start(":1323"))

}

func handleScrape(c echo.Context) error {
	//사용자가 파일을 다운로드 하면
	//서버에서는 파일을 삭제
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrapper(term)
	//첨부파일을 return 하는 기능 즉 서버에서 다운로드 라는 기능 추가됨
	return c.Attachment(fileName, fileName)
}

func handleHome(c echo.Context) error {
	//c.String(http.StatusOK, "Hello, World!") //문자열
	// html 파일을 반환
	return c.File("home.html")
}
