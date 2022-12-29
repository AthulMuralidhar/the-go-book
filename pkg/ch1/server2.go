package ch1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mutexContainer sync.Mutex

func Server2() {
	http.HandleFunc("/", indexHandler2)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/debug", debugHandler)
	http.HandleFunc("/lissajous", lissaJousHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissaJousHandler(writer http.ResponseWriter, request *http.Request) {
	LissaJous1(writer)
}

func debugHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "method: %s\t url: %s\t proto:%s\t\n", request.Method, request.URL, request.Proto)
	for key, value := range request.Header {
		fmt.Fprintf(writer, "header key: %s \t value:%s\n", key, value)
	}
	fmt.Fprintf(writer, "host:%s \n", request.Host)
	fmt.Fprintf(writer, "remote address:%s \n", request.RemoteAddr)

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "error during request.ParseForm: %v", err)
	}

	for key, val := range request.Form {
		fmt.Fprintf(writer, "form key:%s\t value: %s \t", key, val)
	}

}

func countHandler(writer http.ResponseWriter, request *http.Request) {
	mutexContainer.Lock()
	fmt.Fprintf(writer, "counter value:%d", count)
}

func indexHandler2(writer http.ResponseWriter, request *http.Request) {
	mutexContainer.Lock()
	count++
	mutexContainer.Unlock()
	fmt.Fprintf(writer, "given url.path:%s\n", request.URL.Path)
}
