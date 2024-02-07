package game_giant_bomb

import (
	"encoding/json"
	"fmt"

	client_giant_bomb "github.com/aleatoreo22/mygamelist-api/pkg/giant_bomb/client"
	"github.com/aleatoreo22/mygamelist-api/pkg/giant_bomb/model"
)

func GetGameRatings() model.GameRatings {
	response := client_giant_bomb.Get("game_ratings/", nil)
	var gameRatings model.GameRatings
	err := json.Unmarshal([]byte(response), &gameRatings)
	if err != nil {
		fmt.Println(err)
	}
	return gameRatings
}

func GetGame(id string) model.Game {
	response := client_giant_bomb.Get("game_ratings/"+id, nil)
	var game model.Game
	err := json.Unmarshal([]byte(response), &game)
	if err != nil {
		fmt.Println(err)
	}
	return game
}
