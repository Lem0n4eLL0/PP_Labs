package lab2

import (
	"fmt"
	"unicode/utf8"
)

func Init() {
	fmt.Println("\n---------\nLAB2\n---------\n")
	first()
	second()
	third()
	fourth()
	fifth()
	sixth()
}

func first() {
	var num int64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d\n", &num)
	fmt.Println("First task")
	fmt.Print(num, " is ")
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func second() {
	fmt.Println("__________\nSecond task")
	fmt.Println("52 is ", numState(52))
	fmt.Println("-52 is ", numState(-52))
	fmt.Println("0 is ", numState(0))
}

func numState(num float64) string {
	if num < 0 {
		return "Negative"
	} else if num > 0 {
		return "Positive"
	} else {
		return "Zero"
	}

}

func third() {
	fmt.Println("__________\nThird task")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func fourth() {
	fmt.Println("_________\nFourth task")
	fmt.Println("line size 'string' is: ", strLenth("string"))
}

func strLenth(str string) int64 {
	return int64(utf8.RuneCountInString(str))
}

func fifth() {
	fmt.Println("_________\nFifth task")
	var rectangle Rectangle = Rectangle{52, 34}
	fmt.Println("The area of the rectangle")
	rectangle.display()
	fmt.Println("S = ", rectangle.square())
}

func (r Rectangle) square() float64 {
	return r.length * r.width
}

func sixth() {
	fmt.Println("_________\nSixth task")
	fmt.Println("average 3 and 4 is: ", average(3, 4))
}

func average(num1 int64, num2 int64) float64 {
	return (float64(num1) + float64(num2)) / 2
}
func (r Rectangle) display() {
	fmt.Println("length: ", r.length, "\nwidth: ", r.width)
}

type Rectangle struct {
	length float64
	width  float64
}
