package files

import (
	storage "bot/Storage"
	"bot/lid/e"
	"encoding/gob"
	"os"
	"path/filepath"
)

const (
	defaultPerm = 0774
)

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("cant save", err) }()

	//Определяемся куда будет сохранять файл
	filePath := filepath.Join(s.basePath, page.UserName)

	//Все файлы каждого пользователя будем складывать в папку UserName

	//Создаем путь
	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

	//Определимся с названием файла
	fileName, err := fileName(page)
	if err != nil {
		return err
	}

	//Теперь добавляем к пути до файла, имя самого файла
	fPath := filepath.Join(filePath, fileName)

	//Создаем файл
	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(fPath); err != nil {
		return err
	}
	return nil

}

// Для определения названия файла
func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
