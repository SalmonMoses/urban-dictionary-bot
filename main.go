package main

import (
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	urbandict "github.com/davidscholberg/go-urbandict"

	telebot "gopkg.in/tucnak/telebot.v2"
)

var (
	msgTemplate = template.Must(template.ParseFiles("message.txt"))
	strBuilder  = &strings.Builder{}
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		def, err := urbandict.Define(m.Text)
		if err != nil {
			log.Fatal(err)
			bot.Send(m.Sender, err)
		}
		err = msgTemplate.Execute(strBuilder, def)
		if err != nil {
			log.Fatal(err)
			bot.Send(m.Sender, err)
		}
		bot.Send(m.Sender, strBuilder.String())
	})
	bot.Start()
}
