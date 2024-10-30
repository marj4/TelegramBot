package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	//token = flags.Get(token) - должен быть секретным
	token := mustToken()
	fmt.Fprintln(token)

	//tgClient = telegram.New(token) - посредник между кодом и API(BF)

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
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	//Во время выполнения парса в tokene будет лежать ссылка на него(*token)
	flag.Parse()

	if *token == ""{
		log.Fatal("token is incorrect")
	}

	return *token

}