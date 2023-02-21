package main

type Cars struct {
	Brand string
	Model string
	Year int
	Horsepower int
	Color string
	Price float64
	Seats int
	Doors int
	AirConditioning bool
	Consumption struct {
		City float64
		Highway float64
	}
}