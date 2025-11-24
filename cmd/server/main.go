package server

import (
	"log"
	"order-service/internal/config"
	"order-service/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	_, err := db.NewPostgres(cfg.DatabaseURL) //TODO _
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r := gin.Default()

	//dependency injection
	//TODO
	//end dependency

	addr := ":" + cfg.HTTPPort
	log.Println("Listening on " + addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
