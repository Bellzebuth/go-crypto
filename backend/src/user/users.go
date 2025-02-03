package users

import (
	"log"

	"github.com/Bellzebuth/go-crypto/src/db"
	"github.com/Bellzebuth/go-crypto/src/models"
)

func CreateUser(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}

	_, err := db.DB.Model(user).Insert()
	if err != nil {
		log.Println("Erreur lors de l'insertion de l'utilisateur:", err)
		return err
	}
	return nil
}
