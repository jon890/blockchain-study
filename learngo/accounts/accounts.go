package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("can't withdraw")

// Account Struct
type Account struct {
	owner   string
	balance int
}

// NewAccounts creates Account
// 내부적으로 생성한 인스턴스의 주소를 반환함
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// a Account => receiver
// Deposit x amount on your account
// 객체를 호출하면 복사본을 전달한다 (Point Receiver)
// *를 사용하면 복사본을 전달하지 않고 호출한 Object를 변경한다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change Owner of ther account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas : ", a.Balance())
}
