package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"auth-service/app/api/auth"
)

func Controllers(r *gin.Engine, db *sql.DB, redisDb *redis.Client) {

	InitMiddlewares(r, db, redisDb)

	r.POST("/auth/signIn", auth.SignIn)
	r.POST("/auth/refresh", auth.Refresh)
	r.DELETE("/auth/signOut", authMiddleware, auth.SignOut)
	r.DELETE("/auth/signOutAll", authMiddleware, auth.SignOutFromAllDevices)

	r.GET("/auth/accessTokenSecret", auth.GetAccessTokenSecret)
}
