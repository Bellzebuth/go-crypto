package models

type Asset struct {
	Id   string `pg:",pk" json:"id"`
	Name string `json:"name"`
}
