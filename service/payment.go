package service

import (
	"fmt"
	"net/http"

	"github.com/awesome-sphere/as-payment/jwt"
	"github.com/awesome-sphere/as-payment/kafka"
	"github.com/awesome-sphere/as-payment/serializer"
	"github.com/gin-gonic/gin"
)

func TicketPayment(c *gin.Context) {
	is_valid, claim := jwt.AuthorizeToken(c)
	fmt.Printf("%v: %t", claim["user_id"], claim["user_id"])
	if is_valid {
		var payment_s serializer.PaymentSerializer
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
		kafka_message := &kafka.MessageInterface{
			UserID:     int(user_id),
			Price:      payment_s.Price,
			TimeSlotId: payment_s.TimeSlotId,
			TheaterId:  payment_s.TheaterID,
			SeatNumber: payment_s.SeatID,
		}
		is_successful, err := kafka.PushMessage(kafka_message)
		print(is_successful)
		if is_successful {
			c.JSON(http.StatusOK, gin.H{
				"status": "Payment Successful",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed with error",
			"error":  err.Error(),
		})
	}
}