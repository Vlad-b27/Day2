package main

import (
	"ex1/entity"
	"fmt"
)

func main() {
	a := "test"
	c := 2
	d := &a
	f := &c
	fmt.Println(*d, *f)
	passbyValue(c)
	fmt.Println("Value of c after passbyValue", c)
	passbyReference(&c)
	fmt.Println("Value of c after passbyReference", c)
	dog := entity.Dog{
		Name:  "Tom",
		Breed: "German",
	}
	canyouWalk(dog)
	cat := entity.Cat{
		Name: "Tod",
	}
	canyouWalk(cat)
	Talk(cat)
}
func passbyValue(c int) {
	c = 4
}

func passbyReference(c *int) {
	*c = 5
}
func canyouWalk(animal entity.Animal) {
	dog, ok := animal.(entity.Dog)
	if ok {
		fmt.Println("I am called for a dog")
		dog.Bite()
	}
	dog.Bite()
	animal.Walk()
}
func Talk(animal entity.Animal) {

	fmt.Println(animal.Talk("Say my name"))
}
