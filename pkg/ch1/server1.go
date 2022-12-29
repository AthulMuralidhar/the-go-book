package ch1

import (
	"fmt"
	"log"
	"net/http"
)

func Server1() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "given url.path:%s\n", request.URL.Path)
}
