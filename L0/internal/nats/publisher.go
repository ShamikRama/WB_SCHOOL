package nats

import (
	"fmt"
	"log"

	"L0/pkg"
	"github.com/nats-io/stan.go"
)

func LoadMsgToNats(filepath string, clusterID string, clientID string) error {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS: %w", err)
	}
	defer sc.Close()

	msg, err := pkg.ReadJson(filepath)
	if err != nil {
		return fmt.Errorf("failed to read JSON from file: %w", err)
	}

	err = sc.Publish("orders", msg)
	if err != nil {
		log.Println("Error publishing message:", err)
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Println("The message has been sent")
	return nil
}
