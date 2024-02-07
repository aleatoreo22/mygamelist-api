package user_database

import (
	database "github.com/aleatoreo22/mygamelist-api/internal/database/base"
	"github.com/aleatoreo22/mygamelist-api/internal/model"
	"github.com/google/uuid"
)

func Get(id uuid.UUID) (model.User, error) {
	var user model.User
	err := database.Open()
	if err != nil {
		return user, err
	}
	return user, nil
}

func Add(user model.User) error {
	err := database.Open()
	if err == nil {

	}
	return err
}
