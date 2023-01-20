package ch7

import (
	"fmt"
	"log"
	"net/http"
)

type dollars2 float32

func (d dollars2) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database2 map[string]dollars

func (db database2) list(writer http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		_, err := fmt.Fprintf(writer, "%s: %s\n", item, price)
		if err != nil {
			http.Error(writer, fmt.Sprintf("error during Fprintf in /list"), http.StatusBadRequest)
		}
	}
}
func (db database2) price(writer http.ResponseWriter, request *http.Request) {
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

func ServMux1() {
	db := database2{"shoes": 50, "stocks": 3}
	//mux := http.NewServeMux()
	//// both are equivalent http.HandlerFunc <-> mux.HandleFunc
	////mux.Handle("/list", http.HandlerFunc(db.list))
	////mux.Handle("/price", http.HandlerFunc(db.price))
	//
	//mux.HandleFunc("/list", db.list)
	//mux.HandleFunc("/price", db.price)

	// can also work without serve mux like so
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // nil uses the default serve mux
}
