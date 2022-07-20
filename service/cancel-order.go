package service

import (
	"net/http"

	"github.com/awesome-sphere/as-payment/db/models"
	"github.com/awesome-sphere/as-payment/kafka"
	"github.com/awesome-sphere/as-payment/kafka/interfaces"
	"github.com/awesome-sphere/as-payment/serializer"
	"github.com/gin-gonic/gin"
)

func CancelOrder(c *gin.Context) {
	var payment_s serializer.CancelOrderSerializer
	if err := c.BindJSON(&payment_s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	kafka_message := &interfaces.UpdateOrderMessageInterface{
		UserID:     payment_s.UserID,
		TimeSlotId: payment_s.TimeSlotId,
		TheaterId:  payment_s.TheaterID,
		SeatNumber: payment_s.SeatID,
		Status:     string(models.Canceled),
	}
	is_successful, err := kafka.UpdateTopic(kafka_message, kafka.UPDATE_ORDER_TOPIC, payment_s.TheaterID)
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
