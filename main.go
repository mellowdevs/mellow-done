package main

import (
	"fmt"
	"os"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mellowdevs/mellow-done/config"
	"github.com/mellowdevs/mellow-done/routes"
)

func main() {
	godotenv.Load()

	db := config.ConnectDB()

	_, ok_host := os.LookupEnv("HOST")
	port, ok_port := os.LookupEnv("PORT")
	if ok_host && ok_port {

		router := gin.Default()
		router.Use(helmet.Default())

		routes.InitRouter(router)

		router.Run(":" + port)

	} else {
		fmt.Printf("Application runs on port %s\n", port)
	}

	defer db.Close()
}
