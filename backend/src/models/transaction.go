package models

import "time"

type Transaction struct {
	Id             int
	AddressId      int
	Address        int
	Timestamp      time.Time
	Value          int64
	PurchasedPrice int64
}
