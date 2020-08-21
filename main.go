package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var user = &User{}
var storage = &Storage{}

func main() {

	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".

		Token:  "1165371757:AAGy9FFFhOcSYhoB1uPU6JBY_7w4d1_xHvg",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/remindme", func(m *tb.Message) {
		b.Send(m.Sender, m.Payload)
	})

	b.Start()
}
