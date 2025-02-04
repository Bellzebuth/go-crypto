package models

type Asset struct {
	Id   string `pg:",pk"`
	Name string
}
