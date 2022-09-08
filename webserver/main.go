package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"embed"

	"github.com/gorilla/mux"
)

// content holds our static web server content.
//go:embed static
var content embed.FS

func main() {
	PORT := 3000
	s := mux.NewRouter()

	static, _ := fs.Sub(content, "static")
	staticContent := http.FS(static)

	s.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(staticContent))).Methods("GET")
	// s.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./static/index.html")
	// }).Methods("GET")

	s.Handle("/", http.FileServer(staticContent))

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), s)
}
