package main

import (
	"PP_LABS/lab7/date"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"[%s] %s %s %v\n",
			r.Method,
			r.URL.Path,
			r.URL.RawQuery,
			time.Since(startTime),
		)
	})
}

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprint(w, "Unsupported method")
	}
}

func DataHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodPost {
		var data date.DateHandle
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Unsupported json", http.StatusBadRequest)
		}
		data.Display()
		response := map[string]interface{}{
			"message": "Данные получены",
			"data":    data,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Fprint(w, "Unsupported method")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandle)
	router.HandleFunc("/data", DataHandle)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Page")
	})
	http.Handle("/", router)

	loggedMux := loggingMiddleware(router)

	fmt.Println("Server listening...")
	fmt.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}
