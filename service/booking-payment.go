package service

import (
	"net/http"

	"github.com/awesome-sphere/as-payment/kafka"
	"github.com/awesome-sphere/as-payment/kafka/interfaces"
	"github.com/awesome-sphere/as-payment/serializer"
	"github.com/gin-gonic/gin"
)

func AddOrder(c *gin.Context) {
	var payment_s serializer.CreateOrderSerializer
	if err := c.BindJSON(&payment_s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kafka_message := &interfaces.CreateOrderMessageInterface{
		UserID:     payment_s.UserID,
		Price:      payment_s.Price,
		TimeSlotId: payment_s.TimeSlotId,
		TheaterId:  payment_s.TheaterID,
		SeatNumber: payment_s.SeatID,
	}
	is_successful, err := kafka.CreateTopic(kafka_message, kafka.CREATE_ORDER_TOPIC, payment_s.TheaterID)
	print(is_successful)
	if is_successful {
		c.JSON(http.StatusOK, gin.H{
			"status": "Updating Status...",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": "Failed with error",
		"error":  err.Error(),
	})
}
