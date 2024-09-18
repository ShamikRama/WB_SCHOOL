package nats

import (
	"L0/pkg"
	"fmt"
	"log"
)

func LoadMsgToNats(filepath string) {
	sc, err := stan.Connect("test-cluster", "publisher-client")
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()
	msg, err := pkg.ReadJson(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	err = sc.Publish("orders", msg)
	if err != nil {
		log.Fatalln(err)
		fmt.Print("The message has not been sent")
	}
	log.Println("The message has been sent")
}
