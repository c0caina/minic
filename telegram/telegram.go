package telegram

import (
	"log"
	"time"

	"github.com/c0caina/minic/telegram/commands"
	"gopkg.in/telebot.v3"
)

func RunTelegram(token string) {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	{
		b.Handle("/start", commands.Start)
		b.Handle("/replace", commands.Replace)
		b.Handle("/getip", commands.Getip)
	}

	b.Start()
}
