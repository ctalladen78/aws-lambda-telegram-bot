package main

import (
	"github.com/ksopin/aws-lambda-telegram-bot/internal/ginhttp"
)

func main() {
	err := ginhttp.Run()
	if err != nil {
		panic(err)
	}
}
