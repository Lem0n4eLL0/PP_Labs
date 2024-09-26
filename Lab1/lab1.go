package lab1

import (
	"fmt"
	"time"
)

func Init() {
	fmt.Println("\n---------\nLAB1\n---------\n")
	first()
	second()
	third()
	fourth()
	fifth()
	sixth()
}

func first() {
	fmt.Println("First task\n", time.Now())
}

func second() {
	var i int = 0
	var f float64 = 0.0
	var s string = "0"
	var b bool = true
	fmt.Println("__________\nSecond task\nint: ", i, "\nfloat64: ", f, "\nstring: ", s, "\nbool: ", b)
}

func third() {
	i := 0
	f := 0.0
	s := "0"
	b := true
	fmt.Println("__________\nThird task\nint: ", i, "\nfloat64: ", f, "\nstring: ", s, "\nbool: ", b)
}

func fourth() {
	fmt.Println("_________\nFourth task")
	operations(34, 52)
}

func operations(num1 int, num2 int) {
	fmt.Println("num1: ", num1, " | num2: ", num2)
	fmt.Println("sum: ", (num1 + num2))
	fmt.Println("dif: ", (num1 - num2))
	fmt.Println("mult: ", (num1 * num2))
	fmt.Println("priv: ", (num1 / num2))
}

func fifth() {
	fmt.Println("_________\nFifth task")
	fmt.Println("34.34 + 52 = ", oper(34.34, 52, add))
	fmt.Println("34.34 - 52.52 = ", oper(34.34, 52.52, dif))
}

func oper(num1 float64, num2 float64, op func(float64, float64) float64) float64 {
	return op(num1, num2)
}

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func dif(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func sixth() {
	fmt.Println("_________\nSixth task")
	fmt.Println("average: 52 | 34.34 | 76.9\n", average(52, 34.34, 76.9))
}

func average(num1 float64, num2 float64, num3 float64) float64 {
	return (num1 + num2 + num3) / 3
}
