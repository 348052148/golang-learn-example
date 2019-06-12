package models

import "time"

// finnace item
type FinanceItem struct {
	Id            int
	VoucherNumber int
	Date          time.Time
	MatterDesc    string
	DebtorSubject int
	LenderSubjcet int
	LoansType     int
	BalanceType   int
	Balance       float32
}
