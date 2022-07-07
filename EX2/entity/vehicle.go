package entity

import "fmt"

type Vehicle interface {
	Ride()
	NrOfTyres()
	MaximumSpeed()
}
type Bicycle struct {
	Not   int
	Speed int
}
type Car struct {
	Not   int
	Speed int
}
type Plain struct {
	Not   int
	Speed int
}
type Int interface {
	Fly()
}

func (v1 Bicycle) Ride() {
	fmt.Println("You ride a bicycle")
}
func (c1 Car) Ride() {
	fmt.Println("You ride a Car")
}
func (p1 Plain) Ride() {
	fmt.Println("You ride a Plain")
}
func (v1 Bicycle) NrOfTyres() {
	fmt.Println("Number ok tyres:", v1.Not)
}
func (c1 Car) NrOfTyres() {
	fmt.Println("Number ok tyres:", c1.Not)
}
func (p1 Plain) NrOfTyres() {
	fmt.Println("Number ok tyres:", p1.Not)
}
func (v1 Bicycle) MaximumSpeed() {
	fmt.Println("Maximum speed:", v1.Speed)
}
func (c1 Car) MaximumSpeed() {
	fmt.Println("Maximum speed:", c1.Speed)
}
func (p1 Plain) MaximumSpeed() {
	fmt.Println("Maximum speed:", p1.Speed)
}
func (p1 Plain) Fly() {
	fmt.Println("I can fly")
}
