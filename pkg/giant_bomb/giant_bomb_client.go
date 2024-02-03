package giant_bomb

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://www.giantbomb.com/api/"

var token string

func Get(endPoint string, parameters map[string]string) string {
	if parameters == nil {
		parameters = make(map[string]string)
	}
	parameters["api_key"] = token
	parameters["format"] = "json"
	fullUrl := url + endPoint
	fullUrl += "?"
	first := true
	for key, value := range parameters {
		if !first {
			fullUrl += "&"
		}
		fullUrl += strings.Trim(key, " ") + "=" + strings.Trim(value, " ")
		first = false
	}
	response, err := http.Get(fullUrl)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return ""
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return ""
	}
	return string(body)
}
