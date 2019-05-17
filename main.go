package main

import (
	"log"
	"net/http"
)

func main() {
	config := ServiceContainer().GetConfigService().GetConfig()
	err := http.ListenAndServe(config.Host+":"+config.Port, ChiRouter().InitRouter())
	if err != nil {
		log.Fatal(err)
	}
}
