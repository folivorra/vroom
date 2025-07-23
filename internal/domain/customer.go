package domain

type Customer struct {
	Id                  int64  `json:"id"`
	UserId              int64  `json:"user_id"`
	Username            string `json:"username"`
	Phone               string `json:"phone"`
	DriverLicenseNumber string `json:"driver_license_number"`
}
