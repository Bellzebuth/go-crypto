package models

import "time"

type Price struct {
	tableName  struct{} `pg:"prices"`
	Id         int      `pg:",pk"`
	AssetId    string
	Asset      Asset `pg:"rel:has-one"`
	Price      int64
	LastUpdate time.Time
}
