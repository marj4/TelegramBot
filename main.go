package main

import (
	"bot/Clients/telegram"
	"flag"
	"log"
)

func main() {
	//token = flags.Get(token) - должен быть секретным
	token := mustToken()

	//tgClient = telegram.New(token) - собеседник API(BF)
	tgClient := telegram.New(mustHost(), mustToken())

	//fetcher = fetcher.New() - получает события

	//processor = processor.New() - обрабатывает события

	//consumer.Start(fetcher,processor)

}

// Создание функции по получению токена бота
func mustToken() string {
	//получение токена в виде флага
	//Вид к консоли bot -tg-bot-token 'mytoken'
	//в token лежит
	token := flag.String(
		"telegram-bot-token",
		"",
		"token for access to telegram bot",
	)

	//Во время выполнения парса в token будет лежать ссылка на него (*token)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is incorrect")
	}

	return *token

}

func mustHost() string {
	host := flag.String(
		"telegram-bot-host",
		"",
		"host for access to telegram bot",
	)
	if *host == "" {
		log.Fatal("host is incorrect")
	}
}
