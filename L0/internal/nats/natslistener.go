package nats

import (
	"Wildber/internal/database"
	"Wildber/internal/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

func StartListening(DB *database.Database) {
	clusterID := "test-cluster"
	clientID := "your-consumer"
	subject := "channel"
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Error connecting to NATS Streaming: %v", err)
	}
	defer sc.Close()
	// Обработчик сообщений
	messageHandler := func(msg *stan.Msg) {
		var model models.Order
		err = json.Unmarshal(msg.Data, &model)
		if err != nil {
			fmt.Println("Message data Unmarshal error")
			return
		}
		if model.Order_uid == "" {
			fmt.Println("Incorrect data, uid is empty")
			return
		}
		if DB.IsExist(model.Order_uid) {
			fmt.Println("Current record is already in DataBase")
			return
		}
		DB.Insert(model.Order_uid, string(msg.Data))
		fmt.Println("Record is successfully writed in Data Base")
	}

	_, err = sc.Subscribe(subject, messageHandler)
	if err != nil {
		log.Fatalf("Error subscribing to subject: %v", err)
	}

	fmt.Printf("Listening for messages on subject %s...\n", subject)
	select {}
}
