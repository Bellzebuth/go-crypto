package models

import (
	"github.com/Bellzebuth/go-crypto/src/utils"
)

type Transaction struct {
	AddressId         int
	Address           Address `pg:"rel:has-one"`
	PriceId           int
	Price             Price `pg:"rel:has-one"`
	Value             int64
	AvgPurchasedPrice float64

	Gain           float64 `pg:"-"`
	PercentageGain float64 `pg:"-"`
	ActualValue    float64 `pg:"-"`
}

func (a Transaction) ComputeGain() (Transaction, error) {
	value, gain, percentageGain, err := utils.CalculateGain(float64(a.Value), a.AvgPurchasedPrice, a.Price.Price)
	if err != nil {
		return a, err
	}

	a.Gain = gain
	a.PercentageGain = percentageGain
	a.ActualValue = value

	return a, nil
}
