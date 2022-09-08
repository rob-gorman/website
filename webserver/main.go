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
	// var dir string

	// flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	// flag.Parse()
	dir := http.FileServer(http.Dir("./static"))
	s := mux.NewRouter()

	static, _ := fs.Sub(content, "static")
	staticContent := http.FS(static)

	s.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", dir))
	s.PathPrefix("/images/").Handler(http.StripPrefix("/images/", dir))

	// s.Handle("/static/", http.FileServer(staticContent)).Methods("GET")
	// s.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))).Methods("GET")
	// s.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./static/index.html")
	// }).Methods("GET")

	s.Handle("/", http.FileServer(staticContent))

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), s)
}
