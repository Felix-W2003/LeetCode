package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myHandler)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("url:%s,method:%v,header:%v", r.URL.Path, r.Method, r.Header)
	w.Write([]byte("Hello World，Hello go"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	log.Println(username, password)
	w.Write([]byte("Login successful"))
}
