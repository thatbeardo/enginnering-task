package domain

type Car struct {
	Make  string
	Model string
	Year  string
	Price int
	Count int
}

type CarRepository interface {
	GetAllCars() []Car
}
