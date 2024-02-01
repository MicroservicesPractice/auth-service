package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"auth-service/app/api"
	"auth-service/app/config/initializers"
	"auth-service/app/helpers"
)

var SERVER_PORT = helpers.GetEnv("SERVER_PORT")

func init() {
	helpers.CheckRequiredEnvs()

	initializers.InitLogger()
}

func main() {
	redisDb := initializers.ConnectRedis()

	dataBase := initializers.ConnectDb()

	defer dataBase.Close()
	defer redisDb.Close()

	router := gin.Default()

	api.Controllers(router, dataBase, redisDb)

	err := router.Run(fmt.Sprintf(":%v", SERVER_PORT))

	if err != nil {
		log.Panicf("Server listen err: %v", err)
	}

	log.Infof("Server has been started on port %v", SERVER_PORT)
}
