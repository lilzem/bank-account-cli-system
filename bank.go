package main

import "fmt"

type Bank struct {
	accounts map[int]*Account
}

func NewBank() (*Bank) {
	return &Bank{make(map[int]*Account)}
}

func (b *Bank) AddAccount(account *Account) {
	b.accounts[account.Id] = account
}

func (b Bank) Transfer(amount float64, senderId, receiverId int) error {
	sender, senderOk := b.accounts[senderId];
	reciever, recieverOk := b.accounts[receiverId];
	if (!senderOk) {
		return fmt.Errorf("non existing Id - %q", senderId)
	}
	if (!recieverOk) {
		return fmt.Errorf("non existing id - %q", receiverId)
	}
	if (sender.Balance < amount) {
		return fmt.Errorf("insufficient funds: Available - %.2f$; Needed - %.2f$", sender.Balance, amount)
	}
	sender.Withdraw(amount, false);
	reciever.Deposit(amount, false);
	sender.addHistoryRecord(Transfer, -amount, &receiverId)
	reciever.addHistoryRecord(Transfer, amount, &senderId)
	return nil
}

func (b Bank) GetAccount(id int) (*Account, error) {
	acc, ok := b.accounts[id]
	if ok {
		return acc, nil
	}
	return nil, fmt.Errorf("error: no account with this id in the bank")
}

func (b Bank) PrintAccounts() {
	for _, acc := range b.accounts {
		fmt.Printf("Account: %d %s \n", acc.Id, acc.Name)
	}
}