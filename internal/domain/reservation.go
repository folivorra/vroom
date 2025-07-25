package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type ReservationStatus string

const (
	ReservationStatusActive    ReservationStatus = "active"
	ReservationStatusCompleted ReservationStatus = "completed"
	ReservationStatusCancelled ReservationStatus = "cancelled"
)

func (r ReservationStatus) IsValid() bool {
	return r == ReservationStatusActive || r == ReservationStatusCompleted || r == ReservationStatusCancelled
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
	Status        ReservationStatus `json:"status"`
}
