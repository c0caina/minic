package commands

import (
	"github.com/c0caina/minic/minic"
	"gopkg.in/telebot.v3"
)

func Replace(c telebot.Context) error {
	args := map[string]string{"text:": "", "old:": "", "new:": ""}
	parseArg(c.Message().Text, &args)

	if args["text:"] == "" || args["old:"] == "" || args["new:"] == "" {
		return c.Send("Переданы не все аргументы Команда должна иметь следующий вид: `/replace text:<text> old:<text> new:<text>`", telebot.ModeMarkdown)
	}

	return c.Send(minic.Replace(args["text:"], args["old:"], args["new:"]))
}
