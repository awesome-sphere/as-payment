package main

import (
	"github.com/awesome-sphere/as-payment/db"
	"github.com/awesome-sphere/as-payment/jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	// NOTE: Change to ReleaseMode when releasing the app
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// initialze database
	db.InitializeDatabase()
	jwt.InitializeJWTSettings()

	router.Run(":9003")
}
