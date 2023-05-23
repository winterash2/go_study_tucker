package main

import "fmt"

type account struct {
	Balance int
}

func (a *account) WithdrawWithPointer(amount int) {
	a.Balance -= amount
}

func (a account) WithdrawWithValue(amount int) account {
	a.Balance -= amount
	return a
}

func main() {
	var account1 account = account{100}

	account1.WithdrawWithPointer(10)
	fmt.Println(account1.Balance)                       // 90
	account1.WithdrawWithValue(10)
	fmt.Println(account1.Balance)                       // 90
	fmt.Println(account1.WithdrawWithValue(10).Balance) // 80
	fmt.Println(account1.Balance)                       // 90
}
