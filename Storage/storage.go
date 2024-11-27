package storage

import (
	"crypto/sha256"
	"fmt"
	"io"
)

// Методы Storage
type Storage interface {
	Save(p *Page)
	PickRandom(user string) (*Page, error)
	Remove(p *Page) error
	IsExist(p *Page) bool
}

// Создаем тип с которым будет работать Storage
// Страница,которую мы хотим сохранить
type Page struct {
	URl      string
	UserName string
}

func (p Page) Hash() (string, error) {
	//Создаем объект для вычисления хеша
	hash := sha256.New()

	//Считываем ссылку и передаем ее в объект
	if _, err := io.WriteString(hash, p.URl); err != nil {
		return "", err
	}

	if _, err := io.WriteString(hash, p.UserName); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil

}
