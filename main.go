package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/routes"
)

func main() {
	godotenv.Load()

	config.ConnectDB()

	host, ok_host := os.LookupEnv("HOST")
	port, ok_port := os.LookupEnv("PORT")

	if ok_host && ok_port {

		if err := http.ListenAndServe(host+":"+port, routes.Init()); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("HOST or PORT not set in .env")
	}
}
