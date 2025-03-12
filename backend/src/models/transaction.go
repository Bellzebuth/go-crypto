package models

import (
	"github.com/Bellzebuth/go-crypto/src/utils"
)

type Transaction struct {
	Id             string  `pg:",pk" json:"id"`
	AddressId      int     `json:"addressId"`
	Address        Address `pg:"rel:has-one" json:"address"`
	PriceId        int     `json:"priceId"`
	Price          Price   `pg:"rel:has-one" json:"price"`
	TimeStamp      string  `json:"timeStamp"`
	Value          int64   `json:"value"`
	PurchasedPrice float64 `json:"purchasedPrice"`

	Gain           float64 `pg:"-" json:"gain"`
	PercentageGain float64 `pg:"-" json:"percentageGain"`
	ActualValue    float64 `pg:"-" json:"actualValue"`
}

func (a Transaction) ComputeGain() (Transaction, error) {
	value, gain, percentageGain, err := utils.CalculateGain(float64(a.Value), a.PurchasedPrice, a.Price.Price)
	if err != nil {
		return a, err
	}

	a.Gain = gain
	a.PercentageGain = percentageGain
	a.ActualValue = value

	return a, nil
}
