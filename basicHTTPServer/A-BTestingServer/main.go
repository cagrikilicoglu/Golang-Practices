package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	abTestingServer := http.Server{
		Addr: "127.0.0.1:9899",

		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Handler:      &abTestingHandler{},
	}
	log.Fatal(abTestingServer.ListenAndServe())
}

type abTestingHandler struct {
}

func (h *abTestingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	designA := "<html><head><title>The Golang Hotel</title></head><body><p>The Golang Hotel is a relaxing place !</p><p>We offer 20% discount if you call this number : <strong>1234567891</strong></p></body></html>"
	designB := "<html><head><title>The Golang Hotel</title></head><body><h2>The Golang Hotel is a relaxing place !</h2><h5>We offer 20% discount if you call this number<strong>1234567892</strong></h5></body></html>"

	minutes := time.Now().Minute()
	if minutes%2 == 0 {
		w.Write([]byte(designA))
	} else {
		w.Write([]byte(designB))
	}

}
