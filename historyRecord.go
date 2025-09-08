package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type HistoryRecord struct {
	Action Action
	Delta float64
	Timestamp time.Time
	Balance float64
	otherAccId *int
}

func NewHistoryRecord(action Action, delta , balance float64, otherAccId *int) (*HistoryRecord, error) {
	switch action {
	case Withdraw, Deposit, Transfer:
		return &HistoryRecord{Action: action, Delta: delta, Timestamp: time.Now(), Balance: balance, otherAccId: otherAccId}, nil
	default:
		return nil, fmt.Errorf("invalid action: %q", action)
	}
}

func (r *HistoryRecord) ParseDelta() (string, error) {
	switch r.Action {
		case Withdraw:
			return fmt.Sprintf("-$%.2f", r.Delta), nil
		case Deposit:
			return fmt.Sprintf("+$%.2f", r.Delta), nil
		case Transfer:
			if r.Delta > 0 {
				return fmt.Sprintf("+$%.2f", r.Delta), nil
			} else {
				return fmt.Sprintf("-$%.2f", math.Abs(r.Delta)), nil
			}
		default:
			return "", errors.New("error: parsing delta went wrong")
	}

}

func (r *HistoryRecord) Print() {
	delta, err := r.ParseDelta()
	if err != nil {
		fmt.Printf("%q", err)
	}
	if r.otherAccId != nil {
		if r.Delta > 0 {
			fmt.Printf("%q | %q | %q | Balance: $%.2f | From account: %d \n", r.Timestamp.String(), r.Action, delta, r.Balance, *r.otherAccId)
			fmt.Printf("----------\n")
		} else {
			fmt.Printf("%q | %q | %q | Balance: $%.2f | To account: %d \n", r.Timestamp.String(), r.Action, delta, r.Balance, *r.otherAccId)
			fmt.Printf("----------\n")
		}
	} else {
		fmt.Printf("%q | %q | %q | Balance: $%.2f \n", r.Timestamp.String(), r.Action, delta, r.Balance)
		fmt.Printf("----------\n")
	}
}