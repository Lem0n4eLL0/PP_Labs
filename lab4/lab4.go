package lab4

import (
	"fmt"
	"os"
	"strings"
)

func Init() {
	fmt.Println("\n---------\nLAB4\n---------\n")
	first_second_third()
	fourth()
	fifth()
	sixth()
}

func first_second_third() {
	fmt.Println("First task")
	var m = map[string]int{"Vlad": 19, "Kate": 20}
	m["Kostya"] = 20

	displayMap(m)
	fmt.Println("__________\nSecond task")
	fmt.Println("Average age in map: ", averageAge(m))
	fmt.Println("__________\nThird task")
	fmt.Println("Delete 'Vlad' in map is: ", deleteInMap(m, "Vlad"))
	fmt.Println("Delete 'Roma' in map is: ", deleteInMap(m, "Roma"))
	displayMap(m)
}

func averageAge(map1 map[string]int) float64 {
	var sum = 0
	for _, i := range map1 {
		sum += i
	}
	return float64(sum) / float64(len(map1))
}

func deleteInMap(m map[string]int, key string) bool {
	if _, ok := m[key]; ok {
		delete(m, key)
		return true
	} else {
		return false
	}
}

func displayMap(m map[string]int) {
	fmt.Println("map:")
	for name, age := range m {
		fmt.Println("name: ", name, " | age: ", age)
	}
}

func fourth() {
	fmt.Println("_________\nFourth task")
	var str string
	fmt.Println("Enter the string:")
	fmt.Fscan(os.Stdin, &str)
	fmt.Println(str, " to upper case: ", strings.ToUpper(str))
}

func fifth() {
	fmt.Println("_________\nFifth task")
	fmt.Println("Sum is: ", sumNums())

}

func sumNums() float64 {
	fmt.Println("Enter the numbers, the entered 0 will mean the end of the input:")
	var sum, num float64
	for {
		fmt.Fscan(os.Stdin, &num)
		if num == 0 {
			break
		} else {
			sum += num
		}
	}
	return sum
}

func sixth() {
	fmt.Println("_________\nSixth task")
	len := 0
	fmt.Println("Enter the number of int:")
	fmt.Fscan(os.Stdin, &len)
	var arr = make([]int, len)
	for i := 0; i < len; i++ {
		fmt.Fscan(os.Stdin, &arr[i])
	}
	fmt.Printf("arr: %v\n", arr)
	displayReverseArr(arr)
}

func displayReverseArr(arr []int) {
	fmt.Printf("reverse arr: ")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}
