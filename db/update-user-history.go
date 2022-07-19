package db

import (
	"log"
	"time"

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

func UpdateUserHistory(user_id int64, time_slot_id int, theater_id int, seat_number []int, price int64, duration time.Time, status models.OrderStatus) {
	// TO FIX/RECHECK
	err := DB.Create(&models.Order{UserID: user_id, Duration: duration, Price: price, Status: status}).Error
	if err != nil {
		log.Fatalf("Failed to set key: %v", err.Error())
		return
	} else {
		for _, elt := range seat_number {
			err := DB.Create(&models.OrderSeats{SeatID: int64(elt)}).Error
			if err != nil {
				log.Fatalf("Failed to set key: %v", err.Error())
				return
			}
		}
		log.Printf("Successfully updating %d's purchase history", user_id)
	}
}
