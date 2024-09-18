package nats

import (
	"L0/internal/repository"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

func RecieveMsgIntoNats(clusterID string, clientID string) error {
	// Подключение к NATS Streaming
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
	defer sc.Close()

	// Подключение к базе данных
	db, err := repository.ConnectToDB("../../.env")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Подписка на канал "orders"
	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		var orders repository.Order
		if err := json.Unmarshal(m.Data, &orders); err != nil {
			log.Fatalf("Ошибка десериализации сообщения: %v", err)
		}

		// Вставка данных в базу postgres
		// func insertOrder
		if err != nil {
			log.Fatalf("Ошибка вставки данных в базу данных: %v", err)
		}

		log.Println("Сообщение успешно обработано и сохранено в базе данных")
	}, stan.DeliverAllAvailable())

	if err != nil {
		log.Fatalf("Ошибка подписки: %v", err)
	}

	return nil
}

// func insertOrder(db *sql.DB, order *repository.Order){}

// func insertInformationOrder(db *sql.DB, orders *models.Orders){}

// func insertDelivery(db *sql.DB, orders *models.Orders){}

// func insertPayment(db *sql.DB, orders *models.Orders){}

//func insertItems(db *sql.DB, orders *models.Orders){}
