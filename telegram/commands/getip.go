package commands

import (
	"github.com/c0caina/minic/minic"
	"gopkg.in/telebot.v3"
)

func Getip(c telebot.Context) error {
	return c.Send(minic.Getip())
}
