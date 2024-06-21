package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/info", infoHandler)
	mux.HandleFunc("/photo/", photoHandler)
	mux.HandleFunc("/css/", cssHandler)
	mux.HandleFunc("/font/", fontHandler)
	http.ListenAndServe(":8080", mux)
}

type PhotoInfo struct {
	Number   int    `json:"number"`
	Filename string `json:"filename"`
	Location string `json:"location"`
	Season   string `json:"season"`
	Year     int    `json:"year"`
	Camera   string `json:"camera"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := os.ReadFile("photo/info.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	inf := []PhotoInfo{}
	err = json.Unmarshal(f, &inf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, inf)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/info.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := os.ReadFile("photo/info.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	inf := []PhotoInfo{}
	err = json.Unmarshal(f, &inf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, inf)
}

func photoHandler(w http.ResponseWriter, r *http.Request) {
	h := http.StripPrefix("/photo", http.FileServer(http.Dir("./photo/")))
	h.ServeHTTP(w, r)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	h := http.StripPrefix("/css", http.FileServer(http.Dir("./css/")))
	h.ServeHTTP(w, r)
}

func fontHandler(w http.ResponseWriter, r *http.Request) {
	h := http.StripPrefix("/font", http.FileServer(http.Dir("./font/")))
	h.ServeHTTP(w, r)
}
