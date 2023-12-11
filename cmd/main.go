package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"exampleApi/config"
	"exampleApi/db"
	"exampleApi/handlers"
	"exampleApi/helpers"
)

func init() {
	config.InitLogger()
}

func main() {
	redisDb := db.ConnectRedis()

	dataBase := db.ConnectDb()

	defer dataBase.Close()

	router := gin.Default()

	handlers.Handlers(router, dataBase, redisDb)

	var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

	err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}

	log.Infof("Server has been started on port %v", SERVER_PORT)
}
