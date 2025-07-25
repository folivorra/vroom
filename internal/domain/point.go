package domain

import "fmt"

// Point (широта;долгота)
type Point struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

func (p Point) ToWKT() string {
	return fmt.Sprintf("POINT(%f %f)", p.Lat, p.Lng)
}
