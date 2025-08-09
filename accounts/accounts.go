package accounts

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
