package date

import "fmt"

type DateHandle struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewDateHandle(name string, age int) *DateHandle {
	return &DateHandle{name, age}
}
func (d DateHandle) Display() {
	fmt.Println("name: ", d.Name, " age: ", d.Age)
}
