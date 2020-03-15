package telegram

import (
	"log"
	"sync"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// SendMessagesToRooms sends messages to roomIDs
func SendMessagesToRooms(message string, roomIDs []string, telegramBotKey string) {
	bot, err := tb.NewBot(tb.Settings{
		Token:  telegramBotKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalln(err)
	}
	c := make(chan bool)
	var wg sync.WaitGroup

	for _, roomID := range roomIDs {
		chatRoom, err := bot.ChatByID(roomID)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("ID")
		log.Println(chatRoom.Recipient())
		wg.Add(1)
		go func() {
			_, err := bot.Send(chatRoom, message)
			if err != nil {
				log.Fatalln(err)
			}
			c <- true
			wg.Done()
		}()
	}

	wg.Wait()
}
