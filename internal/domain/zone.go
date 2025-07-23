package domain

import (
	"time"
)

type Availability string

const (
	Allowed   Availability = "allowed"
	Forbidden Availability = "forbidden"
	Parking   Availability = "parking"
)

func (a Availability) IsValid() bool {
	return a == Allowed || a == Forbidden || a == Parking
}

type Polygon struct {
	Coordinates [][]Point `json:"coordinates"`
}

type Zone struct {
	Id           int64        `json:"id"`
	Name         string       `json:"name"`
	City         string       `json:"city"`
	BoundaryGeo  Polygon      `json:"boundary_geo"`
	Availability Availability `json:"availability"`
}

type ZoneAssignment struct {
	Id         int64     `json:"id"`
	CarId      int64     `json:"car_id"`
	ZoneId     int64     `json:"zone_id"`
	AssignedAt time.Time `json:"assigned_at"`
}
