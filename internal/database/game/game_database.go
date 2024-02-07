package game_database

import (
	"fmt"

	database "github.com/aleatoreo22/mygamelist-api/internal/database/base"
	"github.com/aleatoreo22/mygamelist-api/internal/model"
	"github.com/google/uuid"
)

func Get(id uuid.UUID) model.Game {
	err := database.Open()
	if err != nil {
		fmt.Println(err)
	}
	var game model.Game
	return game
}
