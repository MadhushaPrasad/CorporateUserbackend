package models

type GeoLocationDetailsModel struct {
	Key       string  `json:"key"`
	Radius    float64 `json:"radius"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
