package domain

type Car struct {
	Make         string
	Model        string
	Year         int
	Price        int
	VehicleCount int `json:"vehicle_count"`
}

type CarRepository interface {
	GetAllCars() ([]Car, error)
}
