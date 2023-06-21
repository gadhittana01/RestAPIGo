package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/giriadhittana01/restapi/domain"
)

func main() {
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		book := domain.Book{
			ID:    "B001",
			Isbn:  "1122",
			Title: "The Book of Me",
			Author: &domain.Author{
				Firstname: "Giri Putra",
				Lastname:  "Adhittana",
			},
		}
		json.NewEncoder(w).Encode(book)
	})

	port := 8000
	log.Println("Serving HTTP on ports :" + strconv.Itoa(port))
	if err := http.ListenAndServe(fmt.Sprintf("%s%d", ":", port), nil); err != nil {
		log.Fatal(err)
	}
}
