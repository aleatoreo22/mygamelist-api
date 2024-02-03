package controller

import (
	"github.com/gin-gonic/gin"
)

func LoadController(controllerModel map[string]func(c *gin.Context), router **gin.Engine) {
	for key, value := range controllerModel {
		(*router).GET(key, value)
	}
}
