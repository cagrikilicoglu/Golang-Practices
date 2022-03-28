package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/cagrikilicoglu/http_errors"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	handlers.AllowedOrigins([]string{"https://example.com"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH"})

	r.Use(loggingMiddleWare)
	r.Use(authenticationMiddleWare)
	s := r.PathPrefix("/products").Subrouter()
	// "/products/{name}"
	s.HandleFunc("/{name}", ProductNameHandler)

	p := r.PathPrefix("/user").Subrouter()
	p.HandleFunc("/", userCreate).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	ShutdownServer(srv, time.Second*10)
}

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func ProductNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//r.URL.Query.Get("param")
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	d := ApiResponse{
		Data: vars["name"],
	}
	resp, _ := json.Marshal(d)
	w.Write(resp)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	var u User
	if r.Header.Get("Content-Type") != "application/json" {
		data, _ := json.Marshal(http_errors.NewRestError(http.StatusBadRequest, http_errors.ContentTypeError.Error()))
		w.Write(data)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		data, _ := json.Marshal(http_errors.NewRestError(http.StatusBadRequest, http_errors.CannotDecodeError.Error()))
		w.Write(data)
		return
	}

	personData, err := json.Marshal(u)
	if err != nil {
		data, _ := json.Marshal(http_errors.NewRestError(http.StatusBadRequest, http_errors.CannotDecodeError.Error()))
		w.Write(data)
		return
	}
	w.Write(personData)
}

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Query())
		next.ServeHTTP(w, r)
	})
}

func authenticationMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/products") {
			token := r.Header.Get("Authorization")
			if token != "" {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "token not found", http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// graceful shutdown

func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

func ErrorResponse(err error) {
	return
}
