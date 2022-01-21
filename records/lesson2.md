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
#### `package learning`은 어떻게 만들지?어떻게 동작하지?
```go
package something

import "fmt"

func sayHello()  {
	fmt.Println("sayHello_A")
}

func SayHello()  {
	fmt.Println("sayHello_B")
}
```


#### 왜 `fmt.Println`의 `Println`은 대문자로 시작하지?
`fmt.Println("Hello world")`
Go의 경우, 만약 function 을 export 하고 싶다면
function 을 대문자로 시작해준다.

```go

package main

import (
	"github.com/gwiyeomgo/go-scrapper/something"
)
func main() {
	something.SayHello()
	//대문자 public ,소문자 private
	//SayHello 는 export 되어진 function 이다.
}
```

#### tip 만약 `fmt pacakge`를 더 알고 싶다면?
`cmd + click`,`ctl + click` 하면 그에 관한 모든 내용을 볼 수 있다.

### Variables and Constants
Go는 Type 언어다
(1)상수
(2)변수
```go

package main

import "fmt"

func main() {
 const name string = "test" //(1)
 //name = "test2"
 //Cannot assign to name
 var age int = 1 //(2)
 age =2
 fmt.Print(name,age)
}
```
(2)의 변수는 `age := 1`럼 타입을 생략해도 go가 자동으로 찾아준다.
이렇게 축약형은 오로지 func 안에서만 가능하고 변수에만 적용 가능하다.


# Function
함수의 형태는 어떻게 만들까?
아래처럼 함수 매개변수의 타입,반환(return) 타입을 정해줘야 작동한다

```go
package main

import "fmt"

func multiply(a int, b int) int {
 return a * b
}
func main() {
 fmt.Println(multiply(2,4))
}
```
`func multiply(a, b int) int `의 경우
go는 a,b 모두 int 타입으로 인식한다.

go의 func 은 여러개의 값을 return 할 수 있다.

```go
package main

import (
	"fmt"
	"strings"
)

func multiply(a int,b int)  int {
	return a * b
}

func lenAndUpper(name string) (int,string){
	return len(name),strings.ToUpper(name)
}
func main() {

	totalLength, _ := lenAndUpper("test")
	fmt.Println(totalLength)
}

```
`lenAndUpper` 함수는 2개의 값을 return 한다
`totalLength, _ := lenAndUpper("test")` 처럼
go는 `_` 을 사용해서 값을 무시 할 수 있다.(컴파일러가 값을 무시한다)

golang은 return 값에 이름을 지정하고
함수 안에서 해당 이름에 값을 할당 할 수 있다.

### [naked return](https://levelup.gitconnected.com/go-naked-returns-4e2094b598e6)
return 할 variable을 굳이 명시하지 않아도 return 가능하다

```go
func lenAndUpper(name string)(length int,upercase string){
	defer fmt.Println("I'm done")
	length=len(name)
	upercase =strings.ToUpper(name)
	return
}
```


### defer
defer는 function이 값을 return하고 나면 실행된다.
자주쓴다.

```go
//recover 할떄 코드
defer func() {
		if r := recover(); r != nil {
			logrus.Errorln("Panic: %v", r)
			os.Exit(1)
		}
	}()
```

```go
defer client.Close()

```