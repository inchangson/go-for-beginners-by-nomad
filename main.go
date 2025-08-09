package main

import (
	"fmt"

	"github.com/inchangson/go-for-beginners-by-nomad/accounts"
)

func main() {
	// private이라 직접 접근 불가, 그렇다고 public으로 둘 수 없음 -> constructor 같은 거 없을까?
	// account := banking.Account{owner: "nico", balance: 1000}
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	fmt.Println(*account)

}
