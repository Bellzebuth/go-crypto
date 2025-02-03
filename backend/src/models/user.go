package models

type User struct {
	ID       int64  `pg:",pk"`
	Username string `pg:",unique"`
	Password string
}
