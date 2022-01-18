
# 1강 go 환경 설정

[repl.it](http://repl.it)

go 설치

[The Go Programming Language](https://go.dev/)

[Tutorial: Get started with Go - The Go Programming Language](https://go.dev/doc/tutorial/getting-started)

## windows

`c:\Program Files\Go`  에 설치

자동적으로 생성되는 Go의 PATH 가 환경변수에 설정됨

`GOPATH` 값 = workspace 폴더(GO PROJECT 폴더)

`GOPATH=C:\Users\go`
`GOROOT=C:\Program Files\Go`

- GOPATH 위치에 go폴더를 생성

`mkdir ~/go`
`cd ~/go`

- go 폴더(GOPATH에 지정된 폴더)에 `bin,pkg,src` 를 생성

— go언어를 사용할 때 자주 쓰는 도구 설치 명령

`go get [golang.org/x/tools/cmd/](http://golang.org/x/tools/cmd/)...`

해당 명령 실행시 아래 1,2 폴더 생성됨

1. bin 폴더 : *.go 소스코드 컴파일을 하면, 실행 가능한 바이너리 파일이 저장된다.
2. pkg 폴더 : 프로젝트에 필요한 패키지가 컴파일 되어, 라이브러리 파일이 저장된다.

`mkdir src`

1. src 폴더 : 사용자가 작성한 소스코드나 사용하려는 오픈소스를 저장하는 곳이다.

---

- 추가 tip

  `mkdir -p $GOPATH/src/github.com/{username}/sample`
  `cd $GOPATH/src/github.com/{username}/sample`

  위와같은 형태로 구조를 잡으면 이후 오픈소스로 공개했을 때 다른 프로그래머들이 go get 명령어로 소스를 쉽게 받아갈 수 있다


---

`github` 에서 새 repository를 생성한다 (`go-scrapper` )

`cd $GOPATH/src/github.com/gwiyeomgo` 에 새로 생성한 `go-scrapper` 을`git clone` 해준다.

`cd go-scrapper`  명령어로 폴더 내부로 이동후

- `touch main.go` 로 main.go 파일을 생성한다
- `code .` 명령어로 go-scrapper 폴더를 열어준다. (vs code인 경우)
- goland 에서 해당 폴더를 열어준다.

### 출처

[예제로 배우는 Go 프로그래밍 - Go 설치와 Go 편집기 소개 (golang.site)](http://golang.site/go/article/2-Go-%EC%84%A4%EC%B9%98%EC%99%80-Go-%ED%8E%B8%EC%A7%91%EA%B8%B0-%EC%86%8C%EA%B0%9C)

[https://grepper.tistory.com/35](https://grepper.tistory.com/35)

[디스커버리 Go 2 - 환경설정 및 시작하기 (amazingguni.github.io)](http://amazingguni.github.io/blog/2016/05/go-chapter-2-1-%EC%84%A4%EC%B9%98_%EB%B0%8F_%EC%8B%9C%EC%9E%91)