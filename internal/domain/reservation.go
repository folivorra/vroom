package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type StatusReservation string

const (
	StatusActive    StatusReservation = "active"
	StatusCompleted StatusReservation = "completed"
	StatusCancelled StatusReservation = "cancelled"
)

func (r StatusReservation) IsValid() bool {
	return r == StatusActive || r == StatusCompleted || r == StatusCancelled
}

type Reservation struct {
	Id            int64             `json:"id"`
	CustomerId    int64             `json:"customer_id"`
	CarId         int64             `json:"car_id"`
	StartAt       time.Time         `json:"start_at"`
	EndAt         time.Time         `json:"end_at"`
	StartLocation Point             `json:"start_location"`
	EndLocation   Point             `json:"end_location"`
	TotalPrice    decimal.Decimal   `json:"total_price"`
	Status        StatusReservation `json:"status"`
}
