/*package nats

import (
	"L0/internal/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func RecieveMsgFromNats(clusterID string, clientID string, cache *repository.Cache, envfile string) error {
	// Подключение к NATS Streaming
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
	defer sc.Close()

	// Подключение к базе данных
	db, err := repository.ConnectToDB(envfile)
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
		if _, exist := cache.Get(orders.OrderUid); !exist { // надо создать новый кэш где-то
			if err = insertOrder(db, &orders); err != nil {
				log.Printf("Ошибка вставки данных в базу данных: %v", err)
				return
			}
			cache.Set(orders.OrderUid, orders)
		}
		log.Println("Сообщение успешно обработано и сохранено в базе данных и в кэше")
	}, stan.DeliverAllAvailable())

	if err != nil {
		log.Fatalf("Ошибка подписки: %v", err)
	}

	return nil
}

func insertOrder(db *sql.DB, orders *repository.Order) error {
	if err := insertInformationOrder(db, orders); err != nil {
		return err
	}

	if err := insertDelivery(db, orders); err != nil {
		return err
	}

	if err := insertPayment(db, orders); err != nil {
		return err
	}

	if err := insertItems(db, orders); err != nil {
		return err
	}
	return nil
}

func insertInformationOrder(db *sql.DB, orders *repository.Order) error {
	_, err := db.Exec(`INSERT INTO information_order(id, order_uid, track_number, entry, local, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
			VALUES (COALESCE((SELECT MAX(id) FROM information_order), 0) + 1, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, orders.OrderUid, orders.TrackNumber, orders.Entry, orders.Local,
		orders.InternalSignature, orders.CustomerId, orders.DeliveryService, orders.Shardkey, orders.SmId, orders.DateCreated, orders.OofShard)
	if err != nil {
		return fmt.Errorf("Ошибка вставки данных в information_order: %v", err)
	}
	return nil
}

func insertDelivery(db *sql.DB, orders *repository.Order) error {
	_, err := db.Exec(`INSERT INTO delivery(id, order_id, name, phone, zip, city, address, region, email)
	VALUES (COALESCE((SELECT MAX(id) FROM delivery), 0) + 1, (SELECT MAX(id) FROM information_order), $1, $2, $3, $4, $5, $6, $7)`, orders.Delivery.Name, orders.Delivery.Phone,
		orders.Delivery.Zip, orders.Delivery.City, orders.Delivery.Address, orders.Delivery.Region, orders.Delivery.Email)
	if err != nil {
		return fmt.Errorf("Ошибка вставки данных в delivery: %v", err)
	}
	return nil
}

func insertPayment(db *sql.DB, orders *repository.Order) error {
	_, err := db.Exec(`INSERT INTO payment(id, order_id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
				VALUES (COALESCE((SELECT MAX(id) FROM payment), 0) + 1, (SELECT MAX(id) FROM information_order), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, orders.Payment.Transaction,
		orders.Payment.RequestId, orders.Payment.Currency, orders.Payment.Provider, orders.Payment.Amount, orders.Payment.PaymentDt, orders.Payment.Bank, orders.Payment.DeliveryCost,
		orders.Payment.GoodsTotal, orders.Payment.CustomFee)
	if err != nil {
		return fmt.Errorf("Ошибка вставки данных в payment: %v", err)
	}
	return nil
}

func insertItems(db *sql.DB, orders *repository.Order) error {
	for _, value := range orders.Items {
		_, err := db.Exec(`INSERT INTO items(id, order_id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
		VALUES (COALESCE((SELECT MAX(id) FROM items), 0) + 1, (SELECT MAX(id) FROM information_order), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, value.ChrtId,
			value.TrackNumber, value.Price, value.Rid, value.Name, value.Sale, value.Size, value.TotalPrice, value.NmID, value.Brand, value.Status)
		if err != nil {
			return fmt.Errorf("Ошибка вставки данных в items: %v", err)
		}
	}
	return nil
}*/

package nats

import (
	"L0/internal/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func RecieveMsgFromNats(clusterID string, clientID string, cache *repository.Cache, envfile string) error {
	// Подключение к NATS Streaming
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к NATS: %v", err)
	}
	defer sc.Close()

	// Подключение к базе данных
	db, err := repository.ConnectToDB(envfile)
	if err != nil {
		return fmt.Errorf("Ошибка подключения к базе данных: %v", err)
	}
	// defer db.Close()

	// Подписка на канал "orders"
	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		var order repository.Order
		if err := json.Unmarshal(m.Data, &order); err != nil {
			log.Printf("Ошибка десериализации сообщения: %v", err)
			return
		}

		// Проверка наличия заказа в кэше
		if _, exists := cache.Get(order.OrderUid); !exists {
			// Вставка данных в базу postgres
			if err := insertOrder(db, &order); err != nil {
				log.Printf("Ошибка вставки данных в базу данных: %v", err)
				return
			}

			// Обновление кэша
			cache.Set(order.OrderUid, order)
		}

		log.Println("Сообщение успешно обработано и сохранено в базе данных и в кэше")
	}, stan.DeliverAllAvailable())

	if err != nil {
		return fmt.Errorf("Ошибка подписки: %v", err)
	}

	return nil
}

func insertOrder(db *sql.DB, order *repository.Order) error {
	if err := insertInformationOrder(db, order); err != nil {
		return err
	}

	if err := insertDelivery(db, order); err != nil {
		return err
	}

	if err := insertPayment(db, order); err != nil {
		return err
	}

	if err := insertItems(db, order); err != nil {
		return err
	}

	return nil
}

func insertInformationOrder(db *sql.DB, order *repository.Order) error {
	_, err := db.Exec(`INSERT INTO information_order(id, order_uid, track_number, entry, local, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES (COALESCE((SELECT MAX(id) FROM information_order), 0) + 1, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, order.OrderUid, order.TrackNumber, order.Entry, order.Local,
		order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard)
	if err != nil {
		return fmt.Errorf("Ошибка вставки данных в information_order: %v", err)
	}
	return nil
}

func insertDelivery(db *sql.DB, order *repository.Order) error {
	_, err := db.Exec(`INSERT INTO delivery(id, order_id, name, phone, zip, city, address, region, email)
    VALUES (COALESCE((SELECT MAX(id) FROM delivery), 0) + 1, (SELECT COALESCE(MAX(id), 0) FROM information_order), $1, $2, $3, $4, $5, $6, $7)`, order.Delivery.Name, order.Delivery.Phone,
		order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return fmt.Errorf("Ошибка вставки данных в delivery: %v", err)
	}
	return nil
}

func insertPayment(db *sql.DB, order *repository.Order) error {
	_, err := db.Exec(`INSERT INTO payment(id, order_id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
        VALUES (COALESCE((SELECT MAX(id) FROM payment), 0) + 1, (SELECT COALESCE(MAX(id), 0) FROM information_order), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, order.Payment.Transaction,
		order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost,
		order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return fmt.Errorf("Ошибка вставки данных в payment: %v", err)
	}
	return nil
}

func insertItems(db *sql.DB, order *repository.Order) error {
	for _, item := range order.Items {
		_, err := db.Exec(`INSERT INTO items(id, order_id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
        VALUES (COALESCE((SELECT MAX(id) FROM items), 0) + 1, (SELECT COALESCE(MAX(id), 0) FROM information_order), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, item.ChrtId,
			item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			return fmt.Errorf("Ошибка вставки данных в items: %v", err)
		}
	}
	return nil
}
