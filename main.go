package main

import (
	"github.com/awesome-sphere/as-payment/db"
	"github.com/awesome-sphere/as-payment/jwt"
	"github.com/awesome-sphere/as-payment/kafka"
	"github.com/awesome-sphere/as-payment/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// NOTE: Change to ReleaseMode when releasing the app
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	// initialze database
	db.InitializeDatabase()
	jwt.InitializeJWTSettings()
	kafka.InitializeKafka()

	router.POST("/payment/submit-payment", service.AddOrder)
	router.POST("/payment/pay-order", service.PayOrder)
	router.Run(":9003")
}
