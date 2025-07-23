package domain

import "fmt"

type StatusCar string

const (
	StatusAvailable StatusCar = "available"
	StatusRented    StatusCar = "rented"
)

func (s StatusCar) IsValid() bool {
	return s == StatusAvailable || s == StatusRented
}

func (p Point) ToWKT() string {
	return fmt.Sprintf("POINT(%f %f)", p.Lat, p.Lng)
}

type Car struct {
	Id         int64     `json:"id"`
	Vin        string    `json:"vin"`
	Year       int       `json:"year"`
	Mileage    int       `json:"mileage"`
	Status     StatusCar `json:"status"`
	CarModelId int64     `json:"car_model_id"`
	Location   Point     `json:"location"`
}

type FuelType string

const (
	FuelTypeGas    FuelType = "gas"
	FuelTypeDiesel FuelType = "diesel"
)

func (ft FuelType) IsValid() bool {
	return ft == FuelTypeGas || ft == FuelTypeDiesel
}

type CarModel struct {
	Id       int64    `json:"id"`
	Brand    string   `json:"brand"`
	Model    string   `json:"model"`
	Class    string   `json:"class"`
	Seats    int      `json:"seats"`
	FuelType FuelType `json:"fuel_type"`
}
