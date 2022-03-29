package models

type Vehicle struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	CheckinTime   string `json:"checkin_time" gorm:"not null"`
	VehicleNumber string `json:"vehicle_number" gorm:"unique; not null"`
	VehicleModel  string `json:"vehicle_model" gorm:"not null"`
	VehicleType   string `json:"vehicle_type" gorm:"not null"`
	OwnerPhone    int    `json:"owner_phone" gorm:"not null"`
}
