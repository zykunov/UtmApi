package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

type AccountSuccess struct {
	Id      uint    `json:"id" default:"1"`
	Balance float64 `json:"balance" default:"500.4"`
	Name    string  `json:"name" default:"John"`
	Surname string  `json:"surname" default:"Doe"`
}

type BalanceSuccess struct {
	Code int    `json:"status_code" default:"200"`
	Meta string `json:"meta" default:"null"`
	Data string `json:"data" default:"400.54"`
}

type BalanceFail struct {
	Code int    `json:"status_code" default:"404"`
	Meta string `json:"meta" default:"null"`
	Data string `json:"data" default:"0"`
}

type SummAmount struct {
	Amount float64 `json:"amount"`
}

func RespondJSON(w *gin.Context, status_code int, data interface{}) {
	log.Println("status code:", status_code)
	var message Message

	message.StatusCode = status_code
	message.Data = data
	w.JSON(200, message)
}
