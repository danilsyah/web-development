package main

import (
	"log"
	"net/http"
)

func main() {
	// instansiasi object route
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler) // routing root
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/profile", profileHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	// jika halaman url tidak ada maka tampilkan 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to my website"))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, Saya sedang belajar web golang"))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hai, nama saya Danil, saya sedang belajar Golang"))
}
