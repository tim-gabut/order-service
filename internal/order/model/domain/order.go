package domain

import "time"

type OrderStatus string

const (
	StatusNew       OrderStatus = "NEW"
	StatusAccepted  OrderStatus = "ACCEPTED"
	StatusInProcess OrderStatus = "IN_PROCESS"
	StatusCompleted OrderStatus = "COMPLETED"
	StatusCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID          int64
	Code        string
	Customer    string
	Status      OrderStatus
	TotalAmount float64
	CreatedAt   time.Time
}
