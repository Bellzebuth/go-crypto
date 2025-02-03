package models

import "time"

type Price struct {
	Id          int
	CryptoId    int
	Crypto      Crypto
	Value       int
	LastUpdated time.Time
}
