package main

import (
	tgClient "example/hello/clients/telegram"
	eventConsumer "example/hello/consumer/event-consumer"
	"example/hello/events/telegram"
	"example/hello/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Printf("service started")
	consumer := eventConsumer.NewConsumer(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

//func mustHost() string {
//	host := flag.String("host", "localhost", "hostname")
//
//	flag.Parse()
//	if *host == "" {
//		log.Fatal("host is not specified")
//	}
//	return *host
//}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
