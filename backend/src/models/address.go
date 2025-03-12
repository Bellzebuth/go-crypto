package models

type Address struct {
	Id           int        `pg:",pk" json:"id"`
	Address      string     `json:"address"`
	Name         string     `json:"name"`
	UserId       int        `json:"userId"`
	User         User       `pg:"rel:has-one" json:"user"`
	BlockchainId int        `json:"blockchainId"`
	Blockchain   Blockchain `pg:"rel:has-one" json:"blockchain"`
}
