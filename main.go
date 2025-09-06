package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action string

const (
	Withdraw Action = "WITHDRAW"
	Deposit Action = "DEPOSIT"
	Transfer Action = "TRANSFER"
)


func main() {
    scanner := bufio.NewScanner(os.Stdin)
    bank := NewBank();
    for {
        fmt.Print("> ")
        
        scanner.Scan()
        input := scanner.Text()
        
        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }
        
        command := parts[0]
        args := parts[1:] 
        
        switch command {
        case "create":
            if err := handleCreate(args, bank); err != nil {
				fmt.Printf("%s \n", err)
			}
        case "deposit":
            if err := handleDeposit(args, bank); err != nil {
				fmt.Printf("%s \n", err)
			}
        case "withdraw":
            if err := handleWithdraw(args, bank); err != nil {
				fmt.Printf("%s \n", err)
			}
        case "history":
            if err := handleHistory(args, bank); err != nil {
				fmt.Printf("%s \n", err)
			}
		case "transfer":
			if err := handleTransfer(args, bank); err != nil {
				fmt.Printf("%s \n", err)
			}
		case "accounts":
			handleAccounts(bank)
		case "balance":
			if err := handleBalance(args, bank); err != nil {
				fmt.Printf("%s", err)
			}
        case "quit", "exit", "q", ":q", ":qa":
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Unknown command:", command)
        }
    }
}

func handleCreate(args []string, bank *Bank) error {
	if len(args) != 2 {
		return fmt.Errorf("error: incorrect input")
	}
	id, err := strconv.Atoi(args[0])
	name := args[1];
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	acc, err := NewAccount(id, name)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	bank.AddAccount(acc);
	fmt.Printf("The account with id: %d has been created and registered in the bank \n", id)
	return nil
}

func handleAccounts(bank *Bank) {
	bank.PrintAccounts()
}

func handleDeposit(args []string, bank *Bank) error {
	if len(args) != 2 {
		return fmt.Errorf("error: incorrect input")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	acc, err := bank.GetAccount(id);
	if err != nil {
		return err
	}
	acc.Deposit(amount, true)
	fmt.Printf("Deposited %.2f$ to account %d. New balance: %.2f$ \n", amount, acc.Id, acc.Balance)
	return nil
}

func handleWithdraw(args []string, bank *Bank) error {
	if len(args) != 2 {
		return fmt.Errorf("error: incorrect input")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	acc, err := bank.GetAccount(id);
	if err != nil {
		return err
	}
	if err := acc.Withdraw(amount, true); err != nil {
		return err
	}
	fmt.Printf("Withdrawn %.2f$ from account %d. New balance: %.2f \n", amount, acc.Id, acc.Balance)
	return nil
}

func handleHistory(args []string, bank *Bank) error {
	if len(args) != 1 {
		return fmt.Errorf("error: incorrect input")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	acc, err := bank.GetAccount(id);
	if err != nil {
		return err
	}
	acc.PrintHistory()
	return nil
}

func handleTransfer(args []string, bank *Bank ) error {
	if len(args) != 3 {
		return fmt.Errorf("error: incorrect input")
	}
	senderId, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	recieverId, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	amount, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	if err := bank.Transfer(amount, senderId, recieverId); err != nil {
		return err
	}
	fmt.Printf("%.2f$ Have been transferred from %d to %d \n", amount, senderId, recieverId)
	return nil
}

func handleBalance(args []string, bank *Bank) error {
	if len(args) != 1 {
		return fmt.Errorf("error: incorrect input")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("error %s", err)
	}
	acc, err := bank.GetAccount(id);
	if err != nil {
		return err
	}
	acc.PrintBalance()
	return nil
}






