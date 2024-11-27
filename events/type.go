package events

//Тут будут описаны 2 интерфейса

type Processor interface {
	Process(e Event) error
}

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Type int

// Создаем список событий, которые буду использовать
const (
	unknown Type = iota
	Message
)

type Event struct {
	Type Type
}
