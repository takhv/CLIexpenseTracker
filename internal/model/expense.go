package model

import "time"

type Expense struct {
	Id          int
	Description string
	CreatedAt   time.Time
	Amount      int
}
