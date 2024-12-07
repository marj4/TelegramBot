package telegram

import (
	"bot/Clients/telegram"
	storage "bot/Storage"
	"bot/events"
)

// Реализоовать оба интерфейса(Fetcher and Processor)
type Processor struct {
	tgClient *telegram.Client
	offset   int
	storage  *storage.Storage
}

func New(client *telegram.Client, storage *storage.Storage) *Processor {
	return &Processor{
		tgClient: client,
		storage:  storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	update, err := p.tgClient.Update(p.offset, limit)
	if err != nil {
		return nil, err
	}

	resEvent := make([]events.Event, 0, len(update))

	for _, upd := range update {
		resEvent = append(resEvent, event(upd))
	}

}

func event(update telegram.Update) events.Event {
	resEvent := events.Event{
		Type:    fetchType(update),
		Message: fetchText(update),
	}
	return resEvent

}

func fetchType(update telegram.Update) events.Type {

}

func fetchText(update telegram.Update) string {
	return update.Message

}
