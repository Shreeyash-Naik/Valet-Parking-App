package api

type CheckinBody struct {
	VehicleNumber string `query:"vehicle_number"`
	VehicleModel  string `query:"vehicle_model"`
	VehicleType   string `query:"vehicle_type"`
	OwnerPhone    int    `query:"owner_phone"`
}

type CheckoutQuery struct {
	ID uint `json:"id"`
}

type SearchQuery struct {
	Query string `json:"query"`
}
