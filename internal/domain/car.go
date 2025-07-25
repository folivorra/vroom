package domain

import "github.com/shopspring/decimal"

type CarStatus string

const (
	CarStatusAvailable CarStatus = "available"
	CarStatusRented    CarStatus = "rented"
)

func (s CarStatus) IsValid() bool {
	return s == CarStatusAvailable || s == CarStatusRented
}

type Car struct {
	Id         int64     `json:"id"`
	Vin        string    `json:"vin"`
	Year       int       `json:"year"`
	Mileage    int       `json:"mileage"`
	Status     CarStatus `json:"status"`
	CarModelId int64     `json:"car_model_id"`
	Location   Point     `json:"location"`
}

type FuelType string

const (
	FuelTypeGas      FuelType = "gas"
	FuelTypeDiesel   FuelType = "diesel"
	FuelTypeElectric FuelType = "electric"
)

func (ft FuelType) IsValid() bool {
	return ft == FuelTypeGas || ft == FuelTypeDiesel || ft == FuelTypeElectric
}

type ClassType string

const (
	ClassTypeEconomy  ClassType = "economy"
	ClassTypeStandard ClassType = "standard"
	ClassTypeComfort  ClassType = "comfort"
	ClassTypeBusiness ClassType = "business"
	ClassTypeSUV      ClassType = "suv"
	ClassTypeCargo    ClassType = "cargo"
)

func (ct ClassType) IsValid() bool {
	return ct == ClassTypeCargo || ct == ClassTypeComfort || ct == ClassTypeStandard || ct == ClassTypeEconomy || ct == ClassTypeBusiness || ct == ClassTypeSUV
}

type CarModel struct {
	Id       int64     `json:"id"`
	Brand    string    `json:"brand"`
	Model    string    `json:"model"`
	Class    ClassType `json:"class"`
	Seats    int       `json:"seats"`
	FuelType FuelType  `json:"fuel_type"`
}

type ClassTariffs struct {
	Id             int64           `json:"id"`
	ClassType      ClassType       `json:"class_type"`
	PricePerMinute decimal.Decimal `json:"price_per_minute"`
}
