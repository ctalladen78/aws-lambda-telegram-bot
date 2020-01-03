package app

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"sync"
)

var (
	instance *App
	iOnce sync.Once
)

type App struct {
	Bot *tgbotapi.BotAPI
}

func New() (*App, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("T_TOKEN"))
	if err != nil {
		return nil, err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//bot.Debug = true

	return &App{Bot: bot}, nil
}

func Get() *App {
	iOnce.Do(func() {
		var err error
		instance, err = New()
		if err != nil {
			panic(err)
		}
	})
	return instance
}



func (a *App) Reply(u *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Reply: " + u.Message.Text)
	a.Bot.Send(msg)
	return nil
}



