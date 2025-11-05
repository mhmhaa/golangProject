package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

// Структура для взаимодействия с данными
type Storage struct {
	key   string
	value string
}

func (s Storage) String() string { return fmt.Sprintf("key=%s, value=%s", s.key, s.value) }

// Функция для сохранения новых записей
func SaveData(buffer bytes.Buffer, storage Storage) error {
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(storage)
	if err != nil {
		return err
	}

	return nil
}

// Функция для удаления записи по ключу
//func (storage *Storage) DeleteData(key string) {
//	delete(storage.data, key)
//}

func main() {
	var buff bytes.Buffer
	storage := Storage{}

	dec := gob.NewDecoder(&buff)
	// Запуск программы без аргументов
	if len(os.Args) == 1 {
		if storage.key == "" {
			fmt.Println("Данные в хранилище отсутствуют")
			return
		}
		decodeErr := dec.Decode(&storage)
		if decodeErr != nil {
			fmt.Printf("Ошибка при выводе данных: %s\n", decodeErr)
			return
		} else {
			fmt.Println("Содержимое хранилища:")
			fmt.Println(storage.String())
			return
		}

	}

	// Запуск программы с аргументами
	args := os.Args[1:]
	operation := args[0]
	storage.key = args[1]

	switch operation {
	case "set":
		if len(args) >= 3 {
			storage.value = args[2]
			fmt.Printf("Добавление данных\n %s: %s\n", storage.key, storage.value)
		} else {
			fmt.Printf("Значение для ключа %s не задано\n", storage.key)
			return
		}
	//case "delete":
	//	_, exists := storage.key
	//	if exists {
	//		storage.DeleteData(storage.key)
	//		fmt.Printf("Удалены данные %s\n", storage.key)
	//	} else {
	//		fmt.Printf("Ключ %s не найден\n", storage.key)
	//		return
	//	}
	default:
		fmt.Printf("Неизвестная операция %s\n", operation)
		return
	}

	// Сохранение данных
	saveErr := SaveData(buff, storage)
	if saveErr != nil {
		fmt.Printf("Ошибка при сохранении данных: %s\n", saveErr)
		return
	} else {
		fmt.Println("Изменения сохранены")
		return
	}
}
