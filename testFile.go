package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

// Понять работу части с работой с файлами
// Изучить методы из видео о написании бота на golang

func main() {
	userName := "Murad"
	fileName := "Файл3"

	if err := Save(userName, fileName); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("File %s is saved!", fileName)
	}

	rFile, err := PickRandom(userName)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Random file:%s", rFile)
	}

}

func Save(userName, fileName string) (err error) {
	//Прописываем путь куда будем сохранять файл
	defaultPath := filepath.Join("C:\\Users\\user\\Desktop\\Users", userName)

	//Создаем директорию по пути defaultPath
	if err := os.MkdirAll(defaultPath, 0774); err != nil {
		return err
	}

	//Получаем путь
	fName, err := hashFileName(fileName)
	if err != nil {
		return err
	}

	//Итоговый путь:"C:\Users\user\Desktop\Users\userName\fPath
	fPath := filepath.Join(defaultPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	//Создаем файл по полному пути fPath

	defer file.Close()

	return nil

}

func PickRandom(user string) (string, error) {
	path := filepath.Join("C:\\Users\\user\\Desktop\\Users", user)

	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "dir is not have files", nil
	}

	index := rand.Intn(len(files))
	file := files[index].Name()

	return file, nil

}

func hashFileName(fileName string) (string, error) {
	hash := md5.New()

	if _, err := io.WriteString(hash, fileName); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
