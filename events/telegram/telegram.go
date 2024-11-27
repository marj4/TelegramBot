package telegram

import "bot/Clients/telegram"

// Реализоовать оба интерфейса(Fetcher and Processor)
type Processor struct {
	tg     *telegram.Client
	offset int
	//storage //хранилище, чтобы сохранять ссылки
}

func New(tg *telegram.Client, offset int) *Processor {}
