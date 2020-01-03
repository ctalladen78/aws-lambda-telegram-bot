package main

import(
	"github.com/ksopin/aws-lambda-telegram-bot/internal/ginlambda"
)

func main() {
	err := ginlambda.Run()
	if err != nil {
		panic(err)
	}
}
