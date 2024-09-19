package main

import (
	"L0/internal/app"
	"L0/internal/nats"
	"L0/internal/repository"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// надо поработать над тем чтобы сделать закрытие базы данных
// надо сделать скрипт чтобы засовывать новый json файл

func main() {
	// Подключение к базе данных
	_, err := repository.ConnectToDB("../.env")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	var f string
	fmt.Print("Вставьте полный путь json файла: ")
	fmt.Scan(&f)
	err = nats.LoadMsgToNats(f, "test-cluster", "subscriber-1")
	if err != nil {
		log.Fatalf("Ошибка загрузки сообщения в NATS: %v", err)
	}

	// Создание кэша
	cache := repository.NewCache()

	// Загрузка данных в кэш
	err = repository.LoadCacheFromDB(cache, "../.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки данных в кэш: %v", err)
	}

	// Запуск сервера
	go app.OrderServer(cache, "../data/htmltemp.html")

	// Подключение к NATS Streaming и обработка сообщений
	err = nats.RecieveMsgFromNats("test-cluster", "subscriber-1", cache, "../.env")
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}

	select {}
}
