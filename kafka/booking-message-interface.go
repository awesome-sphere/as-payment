package kafka

// the one we put into postman

type BookingMessageInterface struct {
	UserID     int   `json:"user_id"`
	TimeSlotId int   `json:"time_slot_id"`
	TheaterId  int   `json:"theater_id"`
	SeatNumber []int `json:"seat_number"`
	Price      int   `json:"price"`
}