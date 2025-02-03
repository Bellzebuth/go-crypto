package models

type User struct {
	tableName struct{} `pg:"users"`
	Id        int64    `pg:",pk"`
	Username  string   `pg:",unique"`
	Password  string
}
