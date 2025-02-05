package models

type Blockchain struct {
	Id   int    `pg:",pk" json:"id"`
	Name string `pg:",unique" json:"name"`
}
