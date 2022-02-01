package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

// error 가 많다면 따로 변수로 만들어서 선언
// 이때 error 이름은 err00 처럼 err 로 시작하도록 작명하여 코드퀄리티를 높힌다
var errNoMoney = errors.New("Cant't withdraw you are poor")

// Function
// NewAccount creates Account
func NewAccount(owner string) *Account {
	//NewAccount 함수는 public 함수이다
	//owner 라는 string 데이터 타입에 매개변수를 갖는다
	//Account 의 주소값을 반환 한다
	account := Account{owner: owner, balance: 0}
	return &account
}

//go receiver
// func 키워드 뒤에 () 안에 receiver를 써준다
//Method
//Method와 function은 동등한 관계이다

//Deposit x amount on your account
/*func (a Account) Deposit(amount int)  {
	//Deposit function안에 a 라는 receiver를 가지고 있다
	// a의 타입은 Account 이다.
	// receiver는 struct 의 첫 글자를 따서 소문자로 지어야 한다 *
	//main에서 Deposit을 호출하면
	// a 는 Account 이지만
	// main 에 선언한 account의 복사본이다
	//실제 account 가 아니다

	a.balance += amount
}*/
func (a *Account) Deposit(amount int) {
	//main 에서 account.Deposit()을 했을 때
	// 실제 account 를 가져오려면
	//리시버의 타입에 *(포인터) 를 붙여준다
	//account를 복사하는 것이 아닌
	//Deposit method를 호출한 account를 사용한다
	//go는 보안을 위해 복사본을 만든다

	a.balance += amount
}
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
	//-라면 출금하면 안되도록
	//error handling 해야 한다
	// nil 은 뭐지? error 는 error 와 nil 있으며 이것은 javascript none 과 같다
}

// ChangeOwner of the account
func (a Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//struct를 복사한 값을 return
//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

//go가 struct 에서 자동적으로 호출해 주는 메서드
func (a Account) String() string {
	//string 으로 return
	return fmt.Sprint("금액: ", a.Balance(), " 계좌 소유주: ", a.Owner())
}
