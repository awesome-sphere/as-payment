package db

import (
	"log"

	"github.com/awesome-sphere/as-payment/models"
)

// order.go
// ID       int64       `json:"id" gorm:"primaryKey;autoincrement;not null"`
// UserID   int64       `json:"user_id" gorm:"not null"`
// Duration time.Time   `json:"duration" gorm:"not null"`
// Price    int64       `json:"price"`
// Status   OrderStatus `json:"order_status" sql:"type:order_status"`

// orderseats.go
// ID      int64 `json:"id" gorm:"primaryKey;autoincrement;not null"`
// SeatID  int64 `json:"seat_id" gorm:"not null"`
// Order   Order `gorm:"foreignKey:OrderID"`
// OrderID int64 `json:"order_id" gorm:"not null"`

func CreateUserHistory(user_id int, time_slot_id int, theater_id int, seat_number []int, price int) {
	event := models.Order{
		UserID:     int64(user_id),
		TimeSlotID: int64(time_slot_id),
		TheaterID:  int64(theater_id),
		Price:      int64(price),
		Status:     models.Awaiting,
	}
	err := DB.Create(&event).Error
	if err != nil {
		log.Fatalf("Failed to update user history: %v", err.Error())
		return
	} else {
		for _, elt := range seat_number {
			history := models.OrderSeats{SeatID: int64(elt), Order: event, OrderID: event.ID}
			err := DB.Create(&history).Error
			if err != nil {
				log.Fatalf("Failed to update booking history: %v", err.Error())
				return
			}
			log.Printf("SeatID: %d | Status: %s | Price: %d", elt, event.Status, event.Price)
		}
		log.Printf("Successfully updating %d's purchase history", user_id)
	}
}
