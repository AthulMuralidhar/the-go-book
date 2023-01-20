package ch7

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	switch request.URL.Path {
	default:
		for item, price := range db {
			_, err := fmt.Fprintf(writer, "%s: %s\n", item, price)
			if err != nil {
				http.Error(writer, fmt.Sprintf("error during Fprintf in /list"), http.StatusBadRequest)
			}
		}
	case "/price":
		item := request.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			writer.WriteHeader(http.StatusNotFound)
			msg := fmt.Sprintf("no such item found: %s\n", request.URL)
			http.Error(writer, msg, http.StatusNotFound)
		}
		_, err := fmt.Fprintf(writer, "price: %s\n", price)
		if err != nil {
			http.Error(writer, fmt.Sprintf("error during Fprintf in /price"), http.StatusBadRequest)
		}
	}

}

func Server1() {
	db := database{"shoes": 50, "stocks": 3}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
