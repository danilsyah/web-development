package main

import (
	"latihan1/handler"
	"log"
	"net/http"
)

func main() {
	// instansiasi object route
	mux := http.NewServeMux()

	// cara lain membuat Handler Function
	aboutHandler := func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("About Page"))
	}

	mux.HandleFunc("/", handler.HomeHandler) // routing root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/profile", handler.ProfileHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/about",aboutHandler)
	// anonymous function
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("halaman test"))
	})


	log.Println("Starting web on port 8081")

	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}