package main

import (
	"ex2/entity"
)

func main() {
	v1 := entity.Bicycle{
		Not:   3,
		Speed: 15,
	}
	c1 := entity.Car{
		Not:   4,
		Speed: 50,
	}
	p1 := entity.Plain{
		Not:   0,
		Speed: 80,
	}
	v1.Ride()
	c1.Ride()
	p1.Ride()
	v1.NrOfTyres()
	c1.NrOfTyres()
	p1.NrOfTyres()
	v1.MaximumSpeed()
	c1.MaximumSpeed()
	p1.MaximumSpeed()
	p1.Fly()
}
