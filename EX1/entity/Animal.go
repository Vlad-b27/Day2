package entity

import "fmt"

type Animal interface {
	Walk()
	Talk(s string) string
}
type WildAnimal interface {
	Bite()
}
type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Walk() {
	fmt.Println("Yes, I acan walk")
}
func (d Dog) Talk(s string) string {
	return fmt.Sprintf("%s baw baw baw", s)
}
func (d Dog) Bite() {
	fmt.Sprintf("Beware I can byte")
}

type Cat struct {
	Name string
}

func (c Cat) Walk() {
	fmt.Println("Yes,I can walk")
}
func (c Cat) Talk(s string) string {
	return fmt.Sprintf("%s My name is: %s", s, c.Name)
}
