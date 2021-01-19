package handler

import (
	"html/template"
	"latihan1/entity"
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
	// data := map[string]interface{}{
	// 	"title" 	: "I'm learning Golang Web",
	// 	"id" 			: 1234,
	// 	"content" : "I'm learning Golang WEB with BWA",
	// }

	// passing struct to view
	// data := entity.Product{ID: 112345, Name: "Honda Jazz", Price: 250000000, Stock: 4}

	// slice of struct
	data := []entity.Product{
		{ID: 1001, Name: "Mobilio",Price: 220000000, Stock: 8},
		{ID: 1002, Name: "Xpander",Price: 270000000, Stock: 2},
		{ID: 1003, Name: "Mobilio",Price: 500000000, Stock: 13},
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

func PostGet(w http.ResponseWriter, r *http.Request){
	method := r.Method

	switch method{
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views/templates","layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is Happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Error is happening, keep calmmm", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		err := r.ParseForm()
		if err != nil{
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name": name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views","result.html"), path.Join("views/templates", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		// w.Write([]byte(name))
		// w.Write( []byte(message))

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}