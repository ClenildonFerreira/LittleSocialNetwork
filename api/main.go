package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println(config.StringConnectionDataBase)

	fmt.Printf("Load in Port: %d", config.Port)
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
