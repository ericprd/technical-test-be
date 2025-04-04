package filter

import "time"

type Filter string

const (
	Username      Filter = "username"
	AccountName   Filter = "account_name"
	AccountNumber Filter = "account_number"
	StartDate     Filter = "start_date"
	EndDate       Filter = "end_date"
)

type FilterSpec struct {
	Username      string
	AccountName   string
	AccountNumber string
	StartDate     time.Time
	EndDate       time.Time
}
