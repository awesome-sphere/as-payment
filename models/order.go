package models

import "time"

type OrderStatus string

const (
	Paid      OrderStatus = "PAID"
	Failed    OrderStatus = "FAILED"
	Awaiting  OrderStatus = "AWAITING"
	Cancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID       int64       `json:"id" gorm:"primaryKey;autoincrement;not null"`
	UserID   int64       `json:"user_id" gorm:"not null"`
	Duration time.Time   `json:"duration" gorm:"not null"`
	Price    int64       `json:"price"`
	Status   OrderStatus `json:"order_status" sql:"type:order_status"`
}
