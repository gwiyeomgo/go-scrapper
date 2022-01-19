### Main Package

만약 프로젝트를 컴파일 하고 싶다면
 main.go 파일이 필요합니다.
```go
package main

```
컴파일러는 
`package main`인 것 부터
찾아내기 때문에
`package main`을 사용한다면
컴파일 하겠다는 것입니다.


목적에 따라 프로젝트 컴파일이 필요가 없을수도 있습니다.
ex) 공유라이브러리,오픈소스 기여
만약 `package learning` 이라고 쓴다면?
당신이 만들고 있는 기능들을 다른 사람들이 사용할 수 있도록 하거나
컴파일 할 수 있도록 할 때 사용합니다.

### main.go 에 function
`package main`만 작성된 main.go를
실행한다면 아래와같은 에러가 발생합니다.
```
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
```
```go
package main

func main() {
//go 프로그램의 시작점
}
```
컴파일러가 자동적으로 `maon package` 와
그 안에 있는 main `function`을 먼저 찾고 실행합니다.

[https://github.com/gocolly/colly](colly github 코드)를
보면 `main.go`가 아닌 `colly.go`파일이 존재한다.
`main.go`는 오직 컴파일을 위해서 필요한 것이다.

### Packages and imports
#### `package learning`은 어떻게 만들지?

#### 왜 `fmt.Println`의 `Println`은 대문자로 시작하지?
`fmt.Println("Hello world")`
Go의 경우, 만약 function 을 export 하고 싶다면
function 을 대문자로 시작해준다.



