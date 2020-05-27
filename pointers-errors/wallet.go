package main

import "fmt"

// Wallet struct
type Wallet struct {
	// lowercase symbol: private outside the package it's defined in
	balance int
}

// Deposit function to know amount user deposit
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance function to know amount of user balance
func (w *Wallet) Balance() int {
	return w.balance
}
