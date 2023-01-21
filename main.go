package main

import (
	"flag"
	"log"

	tgClient "telegram_bot/clients/telegram"
	event_consumer "telegram_bot/consumer/event-consumer"
	"telegram_bot/events/telegram"
	"telegram_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

// 5626315116:AAE6_vCIY8NVzNRFDZb-BqH7L9Br7j6pGOI

func main() {
	// s, err := sqlite.New(sqliteStoragePath)
	// if err != nil {
	// 	log.Fatal("can't connect to storage: ", err)
	// }
	// if err := s.Init(context.TODO()); err != nil {
	// 	log.Fatal("can't init storage", err)
	// }
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
