package main

import "fmt"

// Stringer is implemented by any value that has a String method
type Stringer interface {
	String() string
}

// Bitcoin type (type MyName OriginalType)
type Bitcoin int

// String method for bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	// lowercase symbol: private outside the package it's defined in
	balance Bitcoin
}

// Deposit function to know amount user deposit
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance function to know amount of user balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
