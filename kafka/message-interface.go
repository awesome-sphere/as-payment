package kafka

type MessageInterface struct {
	UserID     int `json:"user_id"`
	TimeSlotId int `json:"time_slot_id"`
	TheaterId  int `json:"theater_id"`
	SeatNumber int `json:"seat_number"`
	SeatTypeId int `json:"seat_type_id"`
	Price      int `json:"price"`
}
