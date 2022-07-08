package models

type Vehicle struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	CheckinTime   string `json:"checkin_time" gorm:"not null"`
	VehicleNumber string `json:"vehicle_number" gorm:"not null"`
	VehicleModel  string `json:"vehicle_model"`
	VehicleType   string `json:"vehicle_type"`
	OwnerPhone    int    `json:"owner_phone"`
}
