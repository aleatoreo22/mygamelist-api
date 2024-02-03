package giant_bomb

import (
	"encoding/json"
	"fmt"
)

func GetGameRatings() GameRatings {
	response := Get("game_ratings/", nil)
	var gameRatings GameRatings
	err := json.Unmarshal([]byte(response), &gameRatings)
	if err != nil {
		fmt.Println(err)
	}
	return gameRatings
}

func LoadToken(tokenParam string) {
	token = tokenParam
}
