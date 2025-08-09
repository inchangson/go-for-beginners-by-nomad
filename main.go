package main

import (
	"fmt"
	"log"

	"github.com/inchangson/go-for-beginners-by-nomad/accounts"
)

func main() {
	// private이라 직접 접근 불가, 그렇다고 public으로 둘 수 없음 -> constructor 같은 거 없을까?
	// account := banking.Account{owner: "nico", balance: 1000}
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	// private field 까지 출력되지만 이건 fmt에서 reflection을 통해 출력되는 것
	// 따라서 외부 시스템 혹은 JSON으로 내보낼 땐 공개 필드만 serialize 됨
	fmt.Println(*account)

	fmt.Println("exec DepositByValueReceiver(100)")
	account.DepositByValueReceiver(100)
	fmt.Println(" -> balance:", account.Balance())

	fmt.Println("exec Deposit(100)")
	account.Deposit(100)
	fmt.Println(" -> balance:", account.Balance())

	fmt.Println("exec Withdraw(50)")
	account.Withdraw(50)
	fmt.Println(" -> balance:", account.Balance())

	fmt.Println("exec Withdraw(1000000000) <- **error 발생**")
	account.Withdraw(1000000000)
	fmt.Println(" -> balance:", account.Balance())
	fmt.Println("**error handling 안 해주면 아무일도 안 일어난다..!**")
	fmt.Println("exec Withdraw(1000000000) <- **error 발생 with handling error**")
	err := account.Withdraw(1000000000)
	if err != nil {
		log.Fatalln(err) // 출력 후 프로그램 종료시켜줌
	}
	fmt.Println("** you can not reach here **")
}
