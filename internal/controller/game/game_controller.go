package game_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetControllers() map[string]func(*gin.Context) {
	return map[string]func(*gin.Context){
		"/game": Get,
	}
}

func Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
