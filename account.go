package main

import (
	"errors"
	"fmt"
)


type Account struct {
	Id int
	Balance float64
	Name string
	History []*HistoryRecord
}

func NewAccount(id int, name string) (*Account, error) {
	if id < 1000 || id > 9999 {
		return nil, errors.New("error: Id must be 4 digits long and have only digits")
	}

	if (len(name) == 0) {
		return nil, errors.New("error: Empty name")
	}
	return &Account{Id: id, Balance: 0.0, Name: name}, nil
}

func (acc *Account) Withdraw(amount float64, addHistory bool) error {
	result := acc.Balance - amount
	if result < 0 {
		return fmt.Errorf("error: Insufficient funds (available: %.2f$)", acc.Balance);
	} else {
		acc.Balance = result
	}
	if (addHistory) {
	acc.addHistoryRecord(Withdraw, amount, nil)
	}
	return nil
}

func (acc *Account) Deposit(amount float64, addHistory bool) {
	acc.Balance += amount
	if (addHistory) {
		acc.addHistoryRecord(Deposit, amount, nil)
	}
}

func (acc *Account) addHistoryRecord(action Action, amount float64, otherAccId *int) {
	record, err := NewHistoryRecord(action, amount, acc.Balance, otherAccId)
	if (err != nil) {
		fmt.Printf("%q", err)
	}
	acc.History = append(acc.History, record)
}

func (acc Account) PrintHistory() {
	for _, record := range acc.History {
		record.Print()
	}
}

func (acc Account) PrintBalance() {
	fmt.Printf("Current balance is: %.2f$ \n", acc.Balance)
}