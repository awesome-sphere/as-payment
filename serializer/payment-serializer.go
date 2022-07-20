package serializer

type CreateOrderSerializer struct {
	TheaterID  int   `json:"theater_id" binding:"required"`
	SeatID     []int `json:"seat_id" binding:"required"`
	TimeSlotId int   `json:"time_slot_id" binding:"required"`
	Price      int   `json:"price" binding:"required"`
}

type UpdateOrderSerializer struct {
	TheaterID  int   `json:"theater_id" binding:"required"`
	SeatID     []int `json:"seat_id" binding:"required"`
	TimeSlotId int   `json:"time_slot_id" binding:"required"`
}
