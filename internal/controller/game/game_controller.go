package game_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	game_database "github.com/aleatoreo22/mygamelist-api/internal/database/game"
)

func GetControllers() map[string]func(*gin.Context) {
	return map[string]func(*gin.Context){
		"/game": Get,
	}
}

func Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
	var guid uuid.UUID
	game_database.Get(guid)
}
