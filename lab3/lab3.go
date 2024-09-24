package lab3

import (
	"PP_LABS/utils/mathutils"
	"PP_LABS/utils/stringutils"
	"fmt"
	"math/rand"
	"unicode/utf8"
)

func Init() {
	fmt.Println("\n---------\nLAB3\n---------\n")
	second()
	third()
	fifth(fourth())
	sixth()
}

func second() {
	fmt.Println("__________\nSecond task:")
	var num int64
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d\n", &num)
	fmt.Println("Factorial ", num, " is: ", mathutils.Factorial(num))
}

func third() {
	fmt.Println("__________\nThird task:")
	fmt.Println("Reverse 'string' is: ", stringutils.Reversestring("string"))
}

func fourth() [5]int {
	fmt.Println("_________\nFourth task:")
	r := rand.New(rand.NewSource(99))
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = r.Int() % 100
	}
	fmt.Printf("arr: %v\n", arr)
	return arr
}

func fifth(arr [5]int) {
	fmt.Println("_________\nFifth task:")
	var slice []int = arr[:]
	fmt.Printf("slice: %v\n", slice)
	slice = append(slice, 52, 43)
	fmt.Printf("slice after add elements: %v\n", slice)
	slice = append(slice[:2], slice[4:]...)
	fmt.Printf("slice after del elements: %v\n", slice)
}

func sixth() {
	fmt.Println("_________\nSixth task:")
	var slice []string = []string{"first", "second", "third", "fifth", "fourth", "sixth"}
	fmt.Printf("In slice: %v\n", slice)
	fmt.Println("Max length string is: " + getMaxLenStr(slice))
}

func getMaxLenStr(slice []string) string {
	var mSum = 0
	var answer string
	for _, str := range slice {
		if utf8.RuneCountInString(str) > mSum {
			mSum = utf8.RuneCountInString(str)
			answer = str
		}
	}
	return answer
}
