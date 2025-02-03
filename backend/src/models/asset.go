package models

type Asset struct {
	tableName struct{} `pg:"assets"`
	Id        string   `pg:",pk"`
	Name      string
}
