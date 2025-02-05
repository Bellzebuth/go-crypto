package models

type Address struct {
	Id           int `pg:",pk"`
	Address      string
	UserId       int
	User         User `pg:"rel:has-one"`
	BlockchainId int
	Blockchain   Blockchain `pg:"rel:has-one"`
}
