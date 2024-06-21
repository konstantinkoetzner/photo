package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/info", infoHandler)
	mux.HandleFunc("/photo/", photoHandler)

	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func photoHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/photo", http.FileServer(http.Dir("./photo/")))
}

type PhotoInfo struct {
	Number   int    `json:"number"`
	Location string `json:"location"`
	Season   string `json:"season"`
	Year     int    `json:"year"`
	Camera   string `json:"camera"`
	FileName string `json:"fileName"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile("photo/info.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	info := []PhotoInfo{}
	err = json.Unmarshal(f, &info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "info.gohtml", info)
}
