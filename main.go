package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Структура для взаимодействия с данными из файла json
type Storage struct {
	data map[string]string
	file string
}

// Создание экземпляра структуры Storage
func NewStorage(file string) *Storage {
	return &Storage{
		data: make(map[string]string),
		file: file,
	}
}

// Функция для загрузки данных из data.json
func (storage *Storage) LoadData() error {
	file, err := os.OpenFile(storage.file, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&storage.data)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// Функция для сохранения новых записей
func (storage *Storage) SaveData() error {
	file, err := os.OpenFile(storage.file, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(storage.data)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// Функция для удаления записи по ключу
func (storage *Storage) DeleteData(key string) {
	delete(storage.data, key)
}


func main() {
	storage := NewStorage("data.json")
	
	// Загрузка данных из файла
	err := storage.LoadData()
	if err != nil {
		fmt.Printf("Ошибка при загрузке данных: %v\n", err)
		return
	}

	// Запуск программы без аргументов
	if len(os.Args) == 1 {
		if len(storage.data) == 0 {
			fmt.Println("Данные в хранилище отсутствуют")
			return
		}

		fmt.Println("Содержимое хранилища:")
		for key, value := range storage.data {
			fmt.Printf("  %s: %s\n", key, value)
		}
		return
	}

	// Запуск программы с аргументами
	args := os.Args[1:]
	operation := args[0]
	key := args[1]

	switch operation {
		case "set":
			value := args[2:]
			if len(value) != 0 {
				storage.data[key] = strings.Join(value, " ")
				fmt.Printf("Добавлены данные\n %s: %s\n", key, value)
			} else {
				fmt.Printf("Значение ключа %s не задано\n", key)
				return
			}
		case "delete":
			_, exists := storage.data[key]
			if exists {
				storage.DeleteData(key)
				fmt.Printf("Удалены данные %s\n", key)
			} else {
				fmt.Printf("Ключ %s не найден\n", key)
				return
			}
		default:
			fmt.Printf("Неизвестная операция %s\n", operation)
			return
	}

	// Сохранение данных
	err = storage.SaveData()
	if err != nil {
		fmt.Printf("Ошибка при сохранении данных: %s\n", err)
		return
	} else {
		fmt.Println("Изменения сохранены")
		return
	}
}