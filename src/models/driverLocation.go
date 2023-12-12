package models

type DriverLocationUpdateModel struct {
	PoolName  string  `json:"pool_name"`
	DriverID  string  `json:"driver_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
