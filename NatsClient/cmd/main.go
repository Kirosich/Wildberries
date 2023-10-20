package main

import (
	"NatsClient/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nats-io/stan.go"
)

func main() {

	clusterID := "test-cluster"
	clientID := "your-publisher"
	subject := "channel"

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Error connecting to NATS Streaming: %v", err)
	}
	defer sc.Close()

	filePath := []string{"./model.json", "./model2.json", "./model3.json", "./modelerr.json"}
	for _, file := range filePath {
		fileData, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Ошибка чтения файла:", err)
			return
		}
		var Data models.Order
		err = json.Unmarshal(fileData, &Data)
		if err != nil {
			fmt.Println("Ошибка распаковки JSON:", err)
			return
		}

		jsonData, err := json.Marshal(Data)
		if err != nil {
			fmt.Println("Ошибка при маршалинге данных:", err)
		}
		message := jsonData
		err = sc.Publish(subject, message)
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}

		fmt.Println("Message sent")
	}
}
