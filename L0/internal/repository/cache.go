package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	data map[string]Order
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]Order),
	}
}

func (c *Cache) Get(key string) (Order, bool) {
	c.RLock()
	defer c.RUnlock()
	order, exists := c.data[key]
	return order, exists
}

func (c *Cache) Set(key string, order Order) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = order
}

func LoadCacheFromDB(order Order) error {
	db, err := ConnectToDB("../../.env")
	if err != nil {
		fmt.Print(err)
		return err
	}
	var id int
	var orders Order
	LoadOrderInformation(db, &order, &id)

	LoadDelivery(db, &order, &id)
	LoadPayment(db, &order, &id)
	LoadItems(db, &order, &id)

	cache.Set(orders.OrderUid, order, 1*time.Hour) // надо где-то написать функцию инициализации кэша, скорее всего в app
	return nil
}

func LoadPayment(db *sql.DB, order *Order, id *int) {
	rows, err := db.Query("SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment WHERE order_id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	rows.Next()

	err = rows.Scan(&order.Payment.Transaction, &order.Payment.RequestId, &order.Payment.Currency, &order.Payment.Provider,
		&order.Payment.Amount, &order.Payment.PaymentDt, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal,
		&order.Payment.CustomFee)

	if err != nil {
		log.Fatal(err)
	}
}

func LoadDelivery(db *sql.DB, order *Order, id *int) {
	rows, err := db.Query("SELECT name, phone, zip, city, address, region, email FROM delivery WHERE order_id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	rows.Next()

	err = rows.Scan(&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City,
		&order.Delivery.Address, &order.Delivery.Region, &order.Delivery.Email)
	if err != nil {
		log.Fatalln(err)
	}
}

func LoadItems(db *sql.DB, order *Order, id *int) {
	var item Item
	rows, err := db.Query("SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE order_id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			log.Fatalln(err)
		}
	}
	order.Items = append(order.Items, item)
}

func LoadOrderInformation(db *sql.DB, order *Order, id *int) {
	rows, err := db.Query("SELECT id, order_uid, track_number, entry, local, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM information_order")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(id, &order.OrderUid, &order.TrackNumber, &order.Entry, &order.Local, &order.InternalSignature,
			&order.CustomerId, &order.DeliveryService, &order.Shardkey, &order.SmId, &order.DateCreated, &order.OofShard)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
