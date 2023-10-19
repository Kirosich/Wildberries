package main

import (
	"Wildber/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, err := config.InitCfg()
	if err != nil {
		log.Fatalf("Error while init config: %v", err)
	}

	fmt.Println("Server is listening on port 8080")
	err = http.ListenAndServe(":8080", config.Router.Routes)
	if err != nil {
		fmt.Println("Error")
	}
}
