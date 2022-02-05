# go-scrapper
쉽고 빠른 Go 시작하기 강의를 듣고,실습하며 scrapper를 만들었습니다.

## lesson 목적
웹사이트에서 데이터를 추출하는 웹 스크래퍼
csv 파일로 만들어서 출력하기😊

Echo go 라이브러리 사용
go의 빠른 속도*

## job scrapper
```
[MAIN]
[getPages]
(1)[getPage] 
    (1)[extractJob]
         .
         .
         .
    (50)[extractJob]
    .
    .
    .
(5)[getPage] 
    (1)[extractJob]
         .
         .
         .
    (50)[extractJob]
```
# goroutine ,channels 적용
```
[MAIN]
[getPages]
[getPage](1) . . .  [getPage](5)
[extractJob](1~50)  [extractJob](1~50)  => 총 250개 동시에         
```
요청정보는 channels를 이용

extractJob 에 gorotine이 종료되면 getPage로 채널을 전달
getPage가 실행이 종료되면 main 함수로 채널을 전송
(getPage와 main 사이에 마지막 goroutine만들기)

# echo를 사용해서 서버 만들기

