package kafka

import "time"

type MessageInterface struct {
	UserID     int       `json:"user_id"`
	TimeSlotId int       `json:"time_slot_id"`
	TheaterId  int       `json:"theater_id"`
	SeatNumber []int     `json:"seat_number"`
	Price      int       `json:"price"`
	Duration   time.Time `json:"duration" gorm:"not null"`
	Status     string    `json:"status"`
}
