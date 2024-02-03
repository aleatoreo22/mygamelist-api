package main

import (
	"fmt"
	"os"
	"strings"

	controller "github.com/aleatoreo22/mygamelist-api/internal/controller/base"
	game_controller "github.com/aleatoreo22/mygamelist-api/internal/controller/game"
	"github.com/aleatoreo22/mygamelist-api/pkg/giant_bomb"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadToken("giant_bomb")
	router := gin.Default()
	controller.LoadController(game_controller.GetControllers(), &router)
	router.Run("localhost:4002")
}

func LoadToken(platform string) {
	file, err := os.Getwd()
	if err != nil {
		fmt.Println("Issue to read file:", err)
		return
	}
	file += "/.env"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Issue to read file:", err)
		return
	}
	tokens := strings.Split(string(content), "\r\n")
	token := ""
	for _, item := range tokens {
		if strings.Contains(item, platform+":") {
			token = strings.ReplaceAll(item, platform+":", "")
			break
		}
	}
	if token == "" {
		fmt.Println("Can't fount token " + platform + "!")
	}
	giant_bomb.LoadToken(token)
}
