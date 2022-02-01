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


# for,range,...args

``` go
func superAdd(numbers ...int) int {
	total := 0
	for _,number := range numbers{
		//for index,number := range numbers{
			total +=number
	}
	return total
}

func main(){
	total := superAdd(1,2,3,4,5,6)
	fmt.Println(total)
}
```

``` go
func canIDrink(age int) bool {
	//variable experssion
	if koreanAge := age +2 ;koreanAge < 18 {
	//if age <18 {
		return false
	}
	return true
}
```

# Switch

``` go
     switch koreanAge := age +2;koreanAge {
	//switch age {
	case 10:
		return false
	case 18:
		return true
	}
```

switch를 사용하면 if-else 난무하는 경우를 피할 수 있다

### Pointers

``` go
func main (){
	a := 2
	b := a
	a = 10
	fmt.Println(a,b) //출력값은 10 2
	fmt.Println(&a,&b) //메모리 주소를 출력한다
}
```

``` go
func main (){
	a := 2
	b := &a
	a = 10
	fmt.Println(&a,b)  //같은 메모리 주소를 출력한다
	fmt.Println(a,*b)  //출력값은 10 10
}
```

``` go
func main (){
	a := 2
	b := &a //a의 주소를 저장
	fmt.Println(*b) // 메모리를 통해 주소의 값을 보고 싶을 떄
	*b = 20 //b는 a의 주소와 연결되어있기 때문에
	fmt.Println(a)  // 출력값은 20
}
```

## Arrays and Slices

go에서 array를 만들때는 크기를 정해줘야 한다.

``` go
	names := [3]string{"a","b"}
	names[2] = "c"
```

array의 크기에 제한없이 요소를 추가 하고 싶을때는 어떻게 하지?
이떄 data type = slice 쓴다
slice는 기본적으로 array 이지만 length가 없다
slice는 item을 추가할때 append()라 불리는 function을 사용한다
append()는 새로운 값이 추가된 slice를 return한다

``` go
	names := []string{"a","b"}
	//append(slice,추가할 item)
	names = append(names,"c")
```

### map

go의 map은 javascript 의 map과 약간 비슷하다
완전히 똑같지는 않다
character 형태의 map을 생성하기 위해서 해야 할 것은

```go
package main

import "fmt"

func main() {
 //map[key]value{} 형태로 data type 써준다
 test := map[string]string{"name": "test", "age": "12"}
 fmt.Print(test)
 //map도 rage를 이용해서 반복문을 이용할 수 있따
 for key,value := range test {
 	fmt.Println(key,value)
 }
}
```


### struct

struct 는 object 와 비슷하면서 map보다 좀 더 유연한게 특징이다

```go
package main

import "fmt"

type person struct {
 name    string
 age     int
 favFood []string
}

func main() {
 favFood := []string{"apple", "ramen"}
 //test := person{"test", 13, favFood}
 //하지만 위 코드처럼 쓴다면 명확하게 보이지 않기 때문에
 //value가 어떤 key값인지 알 수 없다
 //그래서 아래처럼 key 도 같이 써주는 것이 더 명확하다
 test := person {name:"test",age:19,favFood:favFood}
 fmt.Print(test)

}
```
