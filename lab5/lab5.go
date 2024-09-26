package lab5

import (
	"PP_LABS/lab5/book"
	"PP_LABS/lab5/person"
	"PP_LABS/lab5/shapes"
	"fmt"
)

func Init() {
	fmt.Println("\n---------\nLAB5\n---------\n")
	first()
	second()
	third()
	fifth()
	sixth()
}

func first() {
	fmt.Println("First task")
	fmt.Println("Person:")
	var person = person.NewPerson("Vlad", 19)
	person.Display()
}

func second() {
	fmt.Println("__________\nSecond task")
	fmt.Println("Person:")
	var person = person.NewPerson("Vlad", 19)
	person.Display()
	fmt.Println("\nBirthday!")
	person.Birthday()
	person.Display()
}

func third() {
	fmt.Println("__________\nThird task")
	var circle = shapes.NewCircle(12)
	fmt.Println("Circle:")
	circle.Display()
	fmt.Println("Square is: ", circle.Area())
}

func fifth() {
	fmt.Println("_________\nFifth task")
	var shapes = [...]shapes.Shape{shapes.NewCircle(3), shapes.NewCircle(4), shapes.NewReactangle(3, 4)}
	for _, i := range shapes {
		fmt.Println()
		i.Display()
		fmt.Println("area: ", i.Area())
	}
}

func sixth() {
	fmt.Println("_________\nSixth task")
	fmt.Println("New(old) book:")
	var b = book.NewBook("About live", []string{"God"}, []string{"In the beginning was the Word, and the Word was with God, and the Word was God.", "The end"})
	fmt.Println(b.ToString())
}
