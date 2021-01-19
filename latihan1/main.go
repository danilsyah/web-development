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
	// anonymous function
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("halaman test"))
	})
	mux.HandleFunc("/", handler.HomeHandler) // routing root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/profile", handler.ProfileHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/about",aboutHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	// static file css,js,image
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}