package service

import (
	"net/http"

	"github.com/awesome-sphere/as-payment/db/models"
	"github.com/awesome-sphere/as-payment/jwt"
	"github.com/awesome-sphere/as-payment/kafka"
	"github.com/awesome-sphere/as-payment/serializer"
	"github.com/gin-gonic/gin"
)

func PayOrder(c *gin.Context) {
	is_valid, claim := jwt.AuthorizeToken(c)
	if is_valid {
		var payment_s serializer.UpdateOrderSerializer
		if err := c.BindJSON(&payment_s); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user_id, ok := claim["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Unsuccessful to retrieve User ID",
			})
			return
		}
		kafka_message := &kafka.UpdateOrderMessageInterface{
			UserID:     int(user_id),
			TimeSlotId: payment_s.TimeSlotId,
			TheaterId:  payment_s.TheaterID,
			SeatNumber: payment_s.SeatID,
			Status:     string(models.Paid),
		}
		is_successful, err := kafka.UpdateTopic(kafka_message, kafka.UPDATE_ORDER_TOPIC, payment_s.TheaterID)
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
}
