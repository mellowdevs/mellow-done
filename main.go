package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/routes"
)

func main() {
	godotenv.Load()

	db := config.ConnectDB()

	host, ok_host := os.LookupEnv("HOST")
	port, ok_port := os.LookupEnv("PORT")
	if ok_host && ok_port {

		if err := http.ListenAndServe(host+":"+port, routes.Init(db)); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Application runs on port %s\n", port)
		}
	} else {
		log.Fatal("HOST or PORT not set in .env")
	}

	defer db.Close()
}
