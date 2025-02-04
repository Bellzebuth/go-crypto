package models

import "time"

type Price struct {
	Id         int    `pg:",pk"`
	AssetId    string `pg:",unique"`
	Asset      Asset  `pg:"rel:has-one"`
	Price      int64
	LastUpdate time.Time
}
