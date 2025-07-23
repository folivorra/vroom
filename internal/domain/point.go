package domain

// Point (широта;долгота)
type Point struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}
