package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct { // 대문자로 시작해야 외부에서 접근 가능
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (targetAccount Account) DepositByValueReceiver(amount int) {
	// ㄴ reciever: should start with lower case
	fmt.Println("START[DepositByValueReceiver]======")
	fmt.Println(targetAccount)
	targetAccount.balance += amount
	fmt.Println(targetAccount)
	fmt.Println("END  [DepositByValueReceiver]======")
}

// Deposit x amount on your account
func (targetAccount *Account) Deposit(amount int) {
	// ㄴ reciever: should start with lower case
	// ㄴ Go에선 object, struct에 대해서 복사하여 값 전달이 이뤄지기 때문에
	//   주소값으로 전달받아야지 실제 타겟 객체의 값이 바뀐다.
	fmt.Println("START[Deposit]======")
	fmt.Println(targetAccount)
	targetAccount.balance += amount
	fmt.Println(targetAccount)
	fmt.Println("END  [Deposit]======")
}

// Balance of your account
func (a Account) Balance() int {
	//  ㄴ 즉 getter처럼 기존 값 변경 막기 위해선 value receiver로 충분하다
	return a.balance
}

var errNoMoney error = errors.New("sorry, you are poor")

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
