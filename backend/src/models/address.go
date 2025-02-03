package models

type Address struct {
	tableName struct{} `pg:"address"`
	Id        int      `pg:",pk"`
	Address   string
	UserId    int
	User      User `pg:"rel:has-one"`
}
