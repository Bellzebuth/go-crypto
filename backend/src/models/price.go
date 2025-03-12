package models

import "time"

type Price struct {
	Id         int       `pg:",pk" json:"id"`
	AssetId    string    `pg:",unique" json:"assetId"`
	Asset      Asset     `pg:"rel:has-one" json:"asset"`
	Price      int64     `json:"price"`
	LastUpdate time.Time `json:"lastUpdate"`
}
