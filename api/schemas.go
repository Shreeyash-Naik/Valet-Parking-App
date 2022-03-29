package api

type CheckinBody struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleModel  string `json:"vehicle_model"`
	VehicleType   string `json:"vehicle_type"`
	OwnerPhone    int    `json:"owner_phone"`
}

type CheckoutQuery struct {
	ID uint `json:"id"`
}

type SearchQuery struct {
	Query string `json:"query"`
}
