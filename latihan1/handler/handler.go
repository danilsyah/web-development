package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	// jika halaman url tidak ada maka tampilkan 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views/templates","layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Terjadi Kesalahan url", http.StatusInternalServerError)
		return
	}

	// data map , tipe key = string , tipe data value = bebas
	data := map[string]interface{}{
		"title" 	: "I'm learning Golang Web",
		"id" 			: 1234,
		"content" : "I'm learning Golang WEB with BWA",
	}

	// err = tmpl.Execute(w, nil) // tidak mengirim data ke html
	err = tmpl.Execute(w, data) // mengirim data dinamis ke html

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Terjadi Kesalahan url", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, Saya sedang belajar web golang"))
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hai, nama saya Danil, saya sedang belajar Golang"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "Product Page : %d", idNumb)

	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views","product.html"), path.Join("views/templates","layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil{
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}