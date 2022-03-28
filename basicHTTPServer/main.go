package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// create a server
	myServer := &http.Server{
		// set the server address
		Addr: "127.0.0.1:8080",
		// define some specific configuration
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &myHandler{},
	}
	// launch the server
	fmt.Println("starting servers")
	log.Fatal(myServer.ListenAndServe())

}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// toSend := []byte("<html><head></head><body>Hello</body></html>")
	// _, err := w.Write(toSend)
	// if err != nil {
	// 	log.Printf("error while writing on the body %s", err)
	// }
	_, err := os.Open("./tmp/test.json")
	if err != nil {
		log.Println("error file open", err)
	} else {
		log.Println("file opened")
	}

}
