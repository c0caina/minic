package main

import (
	"flag"

	"github.com/c0caina/minic/telegram"
)

func main() {
	//Todo: Нужно реализовать упаковку проекта в Docker и избавится от передачи значений через аргументы командной строки.
	flag.Parse()

	telegram.RunTelegram(flag.Arg(0))
}
