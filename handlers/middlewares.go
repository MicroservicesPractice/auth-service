package handlers

import (
	"database/sql"
	"exampleApi/helpers"
	"exampleApi/helpers/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

func InitMiddlewares(r *gin.Engine, db *sql.DB, redisDb *redis.Client) {
	r.Use(func(c *gin.Context) {
		setDbMiddleware(c, db)
	})

	r.Use(func(c *gin.Context) {
		setRedisDbMiddleware(c, redisDb)
	})

	r.Use(func(c *gin.Context) {
		setStartTime(c)
	})
}

func setDbMiddleware(c *gin.Context, db *sql.DB) {
	c.Set("db", db)
	c.Next()
}

func setRedisDbMiddleware(c *gin.Context, redisDb *redis.Client) {
	c.Set("redis_db", redisDb)
	c.Next()
}

func setStartTime(c *gin.Context) {
	startTime := time.Now()
	c.Set("startTime", startTime)
	c.Next()
}

func authMiddleware(c *gin.Context) {
	_, err := helpers.ParseToken(c.GetHeader("authorization"), helpers.GetEnv("ACCESS_TOKEN_SECRET"))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		log.HttpLog(c, log.Warn, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	// c.Set("userId", claims["user_id"])

	c.Next()
}
