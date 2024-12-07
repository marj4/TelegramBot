package files

import (
	storage "bot/Storage"
	"bot/lid/e"
	"encoding/gob"
	"errors"
	"math/rand"
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
	defer func() { err = e.WrapIfErr("cant save page", err) }()

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
func (s Storage) PickRandom(user string) (page *storage.Page, err error) {
	defer func() { err = e.WrapIfErr("cant pick random page", err) }()

	//Получаем путь к директории куда сохраняются ссылки
	path := filepath.Join(s.basePath, user)

	//Читаем директорию
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	//Проверяем наличие файлов
	if len(files) == 0 {
		return nil, errors.New("no files found")
	}

	//Выбираем случайный файл по индексу
	n := rand.Intn(len(files))
	file := files[n]

	//Строим полный путь дой файла(file)
	filepath := filepath.Join(path, file.Name())

	//Открываем этот путь
	fileOpen, err := os.Open(filepath)
	if err != nil {
		return nil, e.WrapIfErr("cant open file", err)
	}

	fileOpen.Close()

	//Переменая, в которой будет храниться страница
	var p storage.Page

	//Декодируем страницу
	if err := gob.NewDecoder(fileOpen).Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil

}
func (s Storage) Remove(page *storage.Page) (err error) {
	fileName, err := fileName(page)

	if err != nil {
		return e.WrapIfErr("cant remove file", err)
	}

	path := filepath.Join(s.basePath, fileName)

	if err := os.Remove(path); err != nil {
		return e.WrapIfErr("cant remove file", err)
	}
	return nil

}
func (s Storage) IsExist(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, e.WrapIfErr("cant check if file exists", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	if _, err := os.Stat(path); err != nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}

	return false, nil

}

// Для определения названия файла
func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
