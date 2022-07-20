package db

import (
	"log"

	"github.com/awesome-sphere/as-payment/db/models"
)

type orderStruct struct {
	UserID int64  `json:"user_id"`
	Status string `json:"status"`
}

func UpdateUserHistory(UserID, TimeSlotID, TheaterID, SeatID int, status models.OrderStatus) bool {
	var order orderStruct

	tx := DB.Model(&models.Order{}).Where("time_slot_id = ? AND theater_id = ?", TimeSlotID, TheaterID).Joins("JOIN order_seats ON order_seats.order_id = orders.id").Where("order_seats.seat_id = ?", SeatID).First(&models.Order{}).Scan(&order)
	if tx.Error != nil {
		log.Fatalf("Failed to update user history: %v", tx.Error.Error())
		return false
	}

	if (int(order.UserID) == UserID) && (models.OrderStatus(order.Status) == models.Awaiting) {
		log.Printf("This ticket is either not awaiting payment or doesn't belong to this user")
		return false
	}

	tx.Update("status", status)
	if tx.Error != nil {
		log.Fatalf("Failed to update user history: %v", tx.Error.Error())
		return false
	}
	return status == models.Paid
}
