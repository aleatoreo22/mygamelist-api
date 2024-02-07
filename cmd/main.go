package main

import (
	"fmt"
	"os"
	"strings"

	controller "github.com/aleatoreo22/mygamelist-api/internal/controller/base"
	game_controller "github.com/aleatoreo22/mygamelist-api/internal/controller/game"

	"github.com/gin-gonic/gin"
)

func main() {
	/* giant_bomb.LoadToken(LoadEnv("giant_bomb")) */
	LoadEnv("mysql")
	router := gin.Default()
	controller.LoadController(game_controller.GetControllers(), &router)
	router.Run("localhost:4002")
}

func LoadEnv(config string) string {
	file, err := os.Getwd();
	if err != nil {
		fmt.Println("Issue to read file:", err);
		return "";
	}
	file += "/.env"
	content, err := os.ReadFile(file);
	if err != nil {
		fmt.Println("Issue to read file:", err)
		return ""
	}
	tokens := strings.Split(string(content), "\r\n")
	token := ""
	for _, item := range tokens {
		if strings.Contains(item, config+":") {
			token = strings.ReplaceAll(item, config+":", "")
			break
		}
	}
	if token == "" {
		fmt.Println("Can't fount token " + config + "!")
	}
	return token
}
