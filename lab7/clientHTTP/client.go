package main

import (
	"PP_LABS/lab7/date"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	for {
		buff := make([]byte, 1024)
		n, err := resp.Body.Read(buff)
		fmt.Println(string(buff[:n]))
		if n == 0 || err != nil {
			break
		}
	}

	var d date.DateHandle = *date.NewDateHandle("Valera", 78)

	jsonDate, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp1, err := http.Post("http://localhost:8080/data", "application/json", bytes.NewBuffer(jsonDate))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp1.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp1.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ответ сервера:", result)
}
