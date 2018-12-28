package main

import (
	"html"
	"log"
	"net/http"
	"os"
	"time"

	urbandict "github.com/davidscholberg/go-urbandict"

	telebot "gopkg.in/tucnak/telebot.v2"
)

var (
	introduction = "Hello, I am Urban Bot! Send me any English slang and I will define it for you!"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	port := os.Getenv("PORT")

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/start", func(m *telebot.Message) {
		bot.Send(m.Sender, introduction)
	})

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		def, err := urbandict.Define(m.Text)
		if err != nil {
			log.Fatal(err)
			bot.Send(m.Sender, err)
		}
		msg := def.Word + "\n" + def.Definition + "\nExample: " + def.Example
		bot.Send(m.Sender, html.UnescapeString(msg))
	})

	bot.Start()

	http.ListenAndServe(":"+port, nil)
}
