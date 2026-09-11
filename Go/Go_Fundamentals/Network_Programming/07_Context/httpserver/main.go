package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("这是user"))
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.Write([]byte("这是products"))
	})
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("这是orders"))
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
