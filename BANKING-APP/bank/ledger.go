package bank

import (
	"fmt"
)


type LedgerEntry struct {
	BankName string  
	Amount   float64 
}

type Ledger interface{
	AddEntry(bankName string, amount float64)
	GetBalance(bankName string) float64
	PrintLedger()

}

type BankLedger struct { //bank to bank ledger
	entries []LedgerEntry 
}



func NewBankLedger() *BankLedger {
	return &BankLedger{
		entries: []LedgerEntry{},
	}
}

func (l *BankLedger) AddEntry(bankName string, amount float64) {
	l.entries = append(l.entries, LedgerEntry{BankName: bankName, Amount: amount})
}

func (l *BankLedger) GetBalance(bankName string) float64 {
	var balance float64
	for _, entry := range l.entries {
		if entry.BankName == bankName {
			balance += entry.Amount
		}
	}
	return balance
}

func (l *BankLedger) PrintLedger() {
	fmt.Println("Bank Ledger:")
	for _, entry := range l.entries {
		if entry.Amount != 0 {
			fmt.Printf("%s: %.2f\n", entry.BankName, entry.Amount)
		}
	}
}

