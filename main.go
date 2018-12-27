package main

import (
	"log"
	"os"
	"time"

	urbandict "github.com/davidscholberg/go-urbandict"

	telebot "gopkg.in/tucnak/telebot.v2"
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
			bot.Send(m.Sender, err)
		}
		bot.Send(m.Sender, def.Word+"\n"+def.Definition+"\nDefinition by: "+def.Author)
	})
	bot.Start()
}
