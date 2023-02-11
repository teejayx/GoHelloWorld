package main

import (
	"errors"
	"fmt"
)

type Bitcoin float32 

type Stringer interface{
	  String() string
}

type Wallet struct{

	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin){
 w.balance += amount
}


func (w *Wallet) Balance() Bitcoin {
	return w.balance
}


func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
  
	if amount > w.balance {
		return ErrInsufficientFund
	}
	
   w.balance -= amount
   return nil
}

var ErrInsufficientFund = errors.New("Cannot withdraw,insufficient funds")